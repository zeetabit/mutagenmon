package mutagenmon

import (
	"context"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/mutagen-io/mutagen/cmd/mutagen/daemon"
	daemon2 "github.com/mutagen-io/mutagen/pkg/daemon"
	"github.com/mutagen-io/mutagen/pkg/selection"
	serviceSync "github.com/mutagen-io/mutagen/pkg/service/synchronization"
	"github.com/mutagen-io/mutagen/pkg/synchronization"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const IntervalSec = 2

var disconnected = map[synchronization.Status]struct{}{
	synchronization.Status_HaltedOnRootEmptied:    struct{}{},
	synchronization.Status_HaltedOnRootDeletion:   struct{}{},
	synchronization.Status_HaltedOnRootTypeChange: struct{}{},
	synchronization.Status_Disconnected:           struct{}{},
	synchronization.Status_ConnectingBeta:           struct{}{},
}

var watching = map[synchronization.Status]struct{}{
	synchronization.Status_WaitingForRescan:    struct{}{},
	synchronization.Status_Watching:   struct{}{},
}

var syncing = map[synchronization.Status]struct{}{
	synchronization.Status_ConnectingAlpha:    struct{}{},
	synchronization.Status_Scanning:   struct{}{},
	synchronization.Status_Reconciling: struct{}{},
	synchronization.Status_StagingAlpha:           struct{}{},
	synchronization.Status_StagingBeta:           struct{}{},
	synchronization.Status_Transitioning:           struct{}{},
	synchronization.Status_Saving:           struct{}{},
}

type MutagenMon struct {
	menu     map[string]*systray.MenuItem
	states   map[string]*synchronization.State
	daemon   *grpc.ClientConn
	interval time.Duration
	bad      int
	conflict int
	total    int
}

func is(state *synchronization.State, scope map[synchronization.Status]struct{}) bool {
	var ok bool
	_, ok = scope[state.Status]
	return ok
}

func New() (*MutagenMon, error) {
	lock, err := daemon2.AcquireLock()
	if err == nil {
		// should not be here if daemon is running
		err2 := lock.Release()
		if err2 != nil {
			panic(err2)
		}
		return nil, fmt.Errorf("no daemon is running")
	}
	connection, err := daemon.CreateClientConnection(true, false)
	if err != nil {
		return nil, fmt.Errorf("connect to mutagen daemon: %v", err)
	}
	mutagenMon := MutagenMon{
		menu:     map[string]*systray.MenuItem{},
		states:   map[string]*synchronization.State{},
		daemon:   connection,
		interval: IntervalSec * time.Second,
	}
	return &mutagenMon, nil
}

func (self *MutagenMon) SessionStates(ctx context.Context) (map[string]*synchronization.State, error) {
	synchronizationService := serviceSync.NewSynchronizationClient(self.daemon)
	request := &serviceSync.ListRequest{
		Selection: &selection.Selection{All: true},
	}
	response, err := synchronizationService.List(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get list of mutagen sessions: %v", err)
	}
	states := map[string]*synchronization.State{}
	for _, state := range response.SessionStates {
		states[state.Session.Identifier] = state
	}
	return states, nil
}

func (self *MutagenMon) Scheduler() {
	ctx := context.Background()
	ticker := time.NewTicker(self.interval)
	for range ticker.C {
		states, err := self.SessionStates(ctx)
		if err != nil {
			log.Printf("[WARN] get states: %s", err)
		}
		err = self.CheckStates(ctx, states)
		if err != nil {
			log.Printf("[WARN] check states: %s", err)
		}
	}
}

func hasConflicts(state *synchronization.State) bool {
	return len(state.Conflicts) > 0
}

func Icon(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return defaultIcon
	}
	return b
}

func (self *MutagenMon) UpdateMenuItem(item *systray.MenuItem, state *synchronization.State) {
	msg := state.Status.String()
	if len(state.Conflicts) == 1 {
		msg += fmt.Sprintf(", %d conflict", len(state.Conflicts))
	}
	if len(state.Conflicts) > 1 {
		msg += fmt.Sprintf(", %d conflicts", len(state.Conflicts))
	}
	item.SetTooltip(msg)
	log.Printf("[INFO] state changed to: %s", msg)

	if is(state, disconnected) {
		item.SetIcon(Icon("disconnected.png"))
	} else if is(state, syncing) {
		SetIfNoConflict(state, item, "syncing.png")
	} else if is(state, watching) {
		SetIfNoConflict(state, item, "ok.png")
	} else {
		SetIfNoConflict(state, item, "unknown.png")
	}
}

func SetIfNoConflict(state *synchronization.State, item *systray.MenuItem, name string) {
	if len(state.Conflicts) > 0 {
		item.SetIcon(Icon("conflict.png"))
		return
	}
	item.SetIcon(Icon(name))
}

func (self *MutagenMon) CheckStates(ctx context.Context, states map[string]*synchronization.State) error {
	var bad int
	var conflict int
	for id, current := range states {
		old, ok := self.states[id]
		if !ok {
			item := systray.AddMenuItem(fmt.Sprintf("%s:%s", current.Session.Beta.Host, current.Session.Beta.Path), "")
			self.menu[id] = item
			self.UpdateMenuItem(self.menu[id], current)
			self.states[id] = current
			if is(current, disconnected) {
				bad++
			}
			if hasConflicts(current) {
				conflict++
			}
			continue
		}
		if old.Status != current.Status || len(old.Conflicts) != len(current.Conflicts) {
			self.UpdateMenuItem(self.menu[id], current)
			self.states[id] = current
		}
		if is(current, disconnected) {
			bad++
		}
		if hasConflicts(current) {
			conflict++
		}
	}
	for id, _ := range self.states {
		_, ok := states[id]
		if !ok {
			self.menu[id].Hide()
			delete(self.menu, id)
			delete(self.states, id)
		}
	}
	total := len(self.states)
	if bad != self.bad || total != self.total || conflict != self.conflict {
		systray.SetTitle(fmt.Sprintf("%d / %d / %d", total-conflict - bad, total-bad, total))
	}
	self.bad = bad
	self.conflict = conflict
	self.total = total
	log.Printf("[DEBUG] states checked, goroutines: %d", runtime.NumGoroutine())
	return nil
}

func (self *MutagenMon) Run() {
	ep, err := os.Executable()
	if err != nil {
		log.Fatalln("os.Executable:", err)
	}
	err = os.Chdir(filepath.Join(filepath.Dir(ep), "..", "Resources"))
	if err != nil {
		log.Fatalln("os.Chdir:", err)
	}
	systray.Run(self.Init, nil)
}

func (self *MutagenMon) Init() {
	systray.SetTooltip("no conflict / connected / total")
	mQuit := systray.AddMenuItem("Quit", "")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
	systray.AddSeparator()
	go self.Scheduler()
}
