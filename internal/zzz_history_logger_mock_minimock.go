// Code generated by http://github.com/gojuno/minimock (v3.4.0). DO NOT EDIT.

package internal

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// HistoryLoggerMock implements historyLogger
type HistoryLoggerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcLog          func(s *Shipment) (err error)
	funcLogOrigin    string
	inspectFuncLog   func(s *Shipment)
	afterLogCounter  uint64
	beforeLogCounter uint64
	LogMock          mHistoryLoggerMockLog
}

// NewHistoryLoggerMock returns a mock for historyLogger
func NewHistoryLoggerMock(t minimock.Tester) *HistoryLoggerMock {
	m := &HistoryLoggerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.LogMock = mHistoryLoggerMockLog{mock: m}
	m.LogMock.callArgs = []*HistoryLoggerMockLogParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mHistoryLoggerMockLog struct {
	optional           bool
	mock               *HistoryLoggerMock
	defaultExpectation *HistoryLoggerMockLogExpectation
	expectations       []*HistoryLoggerMockLogExpectation

	callArgs []*HistoryLoggerMockLogParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// HistoryLoggerMockLogExpectation specifies expectation struct of the historyLogger.Log
type HistoryLoggerMockLogExpectation struct {
	mock               *HistoryLoggerMock
	params             *HistoryLoggerMockLogParams
	paramPtrs          *HistoryLoggerMockLogParamPtrs
	expectationOrigins HistoryLoggerMockLogExpectationOrigins
	results            *HistoryLoggerMockLogResults
	returnOrigin       string
	Counter            uint64
}

// HistoryLoggerMockLogParams contains parameters of the historyLogger.Log
type HistoryLoggerMockLogParams struct {
	s *Shipment
}

// HistoryLoggerMockLogParamPtrs contains pointers to parameters of the historyLogger.Log
type HistoryLoggerMockLogParamPtrs struct {
	s **Shipment
}

// HistoryLoggerMockLogResults contains results of the historyLogger.Log
type HistoryLoggerMockLogResults struct {
	err error
}

// HistoryLoggerMockLogOrigins contains origins of expectations of the historyLogger.Log
type HistoryLoggerMockLogExpectationOrigins struct {
	origin  string
	originS string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmLog *mHistoryLoggerMockLog) Optional() *mHistoryLoggerMockLog {
	mmLog.optional = true
	return mmLog
}

// Expect sets up expected params for historyLogger.Log
func (mmLog *mHistoryLoggerMockLog) Expect(s *Shipment) *mHistoryLoggerMockLog {
	if mmLog.mock.funcLog != nil {
		mmLog.mock.t.Fatalf("HistoryLoggerMock.Log mock is already set by Set")
	}

	if mmLog.defaultExpectation == nil {
		mmLog.defaultExpectation = &HistoryLoggerMockLogExpectation{}
	}

	if mmLog.defaultExpectation.paramPtrs != nil {
		mmLog.mock.t.Fatalf("HistoryLoggerMock.Log mock is already set by ExpectParams functions")
	}

	mmLog.defaultExpectation.params = &HistoryLoggerMockLogParams{s}
	mmLog.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmLog.expectations {
		if minimock.Equal(e.params, mmLog.defaultExpectation.params) {
			mmLog.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLog.defaultExpectation.params)
		}
	}

	return mmLog
}

// ExpectSParam1 sets up expected param s for historyLogger.Log
func (mmLog *mHistoryLoggerMockLog) ExpectSParam1(s *Shipment) *mHistoryLoggerMockLog {
	if mmLog.mock.funcLog != nil {
		mmLog.mock.t.Fatalf("HistoryLoggerMock.Log mock is already set by Set")
	}

	if mmLog.defaultExpectation == nil {
		mmLog.defaultExpectation = &HistoryLoggerMockLogExpectation{}
	}

	if mmLog.defaultExpectation.params != nil {
		mmLog.mock.t.Fatalf("HistoryLoggerMock.Log mock is already set by Expect")
	}

	if mmLog.defaultExpectation.paramPtrs == nil {
		mmLog.defaultExpectation.paramPtrs = &HistoryLoggerMockLogParamPtrs{}
	}
	mmLog.defaultExpectation.paramPtrs.s = &s
	mmLog.defaultExpectation.expectationOrigins.originS = minimock.CallerInfo(1)

	return mmLog
}

// Inspect accepts an inspector function that has same arguments as the historyLogger.Log
func (mmLog *mHistoryLoggerMockLog) Inspect(f func(s *Shipment)) *mHistoryLoggerMockLog {
	if mmLog.mock.inspectFuncLog != nil {
		mmLog.mock.t.Fatalf("Inspect function is already set for HistoryLoggerMock.Log")
	}

	mmLog.mock.inspectFuncLog = f

	return mmLog
}

// Return sets up results that will be returned by historyLogger.Log
func (mmLog *mHistoryLoggerMockLog) Return(err error) *HistoryLoggerMock {
	if mmLog.mock.funcLog != nil {
		mmLog.mock.t.Fatalf("HistoryLoggerMock.Log mock is already set by Set")
	}

	if mmLog.defaultExpectation == nil {
		mmLog.defaultExpectation = &HistoryLoggerMockLogExpectation{mock: mmLog.mock}
	}
	mmLog.defaultExpectation.results = &HistoryLoggerMockLogResults{err}
	mmLog.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmLog.mock
}

// Set uses given function f to mock the historyLogger.Log method
func (mmLog *mHistoryLoggerMockLog) Set(f func(s *Shipment) (err error)) *HistoryLoggerMock {
	if mmLog.defaultExpectation != nil {
		mmLog.mock.t.Fatalf("Default expectation is already set for the historyLogger.Log method")
	}

	if len(mmLog.expectations) > 0 {
		mmLog.mock.t.Fatalf("Some expectations are already set for the historyLogger.Log method")
	}

	mmLog.mock.funcLog = f
	mmLog.mock.funcLogOrigin = minimock.CallerInfo(1)
	return mmLog.mock
}

// When sets expectation for the historyLogger.Log which will trigger the result defined by the following
// Then helper
func (mmLog *mHistoryLoggerMockLog) When(s *Shipment) *HistoryLoggerMockLogExpectation {
	if mmLog.mock.funcLog != nil {
		mmLog.mock.t.Fatalf("HistoryLoggerMock.Log mock is already set by Set")
	}

	expectation := &HistoryLoggerMockLogExpectation{
		mock:               mmLog.mock,
		params:             &HistoryLoggerMockLogParams{s},
		expectationOrigins: HistoryLoggerMockLogExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmLog.expectations = append(mmLog.expectations, expectation)
	return expectation
}

// Then sets up historyLogger.Log return parameters for the expectation previously defined by the When method
func (e *HistoryLoggerMockLogExpectation) Then(err error) *HistoryLoggerMock {
	e.results = &HistoryLoggerMockLogResults{err}
	return e.mock
}

// Times sets number of times historyLogger.Log should be invoked
func (mmLog *mHistoryLoggerMockLog) Times(n uint64) *mHistoryLoggerMockLog {
	if n == 0 {
		mmLog.mock.t.Fatalf("Times of HistoryLoggerMock.Log mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmLog.expectedInvocations, n)
	mmLog.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmLog
}

func (mmLog *mHistoryLoggerMockLog) invocationsDone() bool {
	if len(mmLog.expectations) == 0 && mmLog.defaultExpectation == nil && mmLog.mock.funcLog == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmLog.mock.afterLogCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmLog.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Log implements historyLogger
func (mmLog *HistoryLoggerMock) Log(s *Shipment) (err error) {
	mm_atomic.AddUint64(&mmLog.beforeLogCounter, 1)
	defer mm_atomic.AddUint64(&mmLog.afterLogCounter, 1)

	mmLog.t.Helper()

	if mmLog.inspectFuncLog != nil {
		mmLog.inspectFuncLog(s)
	}

	mm_params := HistoryLoggerMockLogParams{s}

	// Record call args
	mmLog.LogMock.mutex.Lock()
	mmLog.LogMock.callArgs = append(mmLog.LogMock.callArgs, &mm_params)
	mmLog.LogMock.mutex.Unlock()

	for _, e := range mmLog.LogMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmLog.LogMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLog.LogMock.defaultExpectation.Counter, 1)
		mm_want := mmLog.LogMock.defaultExpectation.params
		mm_want_ptrs := mmLog.LogMock.defaultExpectation.paramPtrs

		mm_got := HistoryLoggerMockLogParams{s}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.s != nil && !minimock.Equal(*mm_want_ptrs.s, mm_got.s) {
				mmLog.t.Errorf("HistoryLoggerMock.Log got unexpected parameter s, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmLog.LogMock.defaultExpectation.expectationOrigins.originS, *mm_want_ptrs.s, mm_got.s, minimock.Diff(*mm_want_ptrs.s, mm_got.s))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLog.t.Errorf("HistoryLoggerMock.Log got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmLog.LogMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLog.LogMock.defaultExpectation.results
		if mm_results == nil {
			mmLog.t.Fatal("No results are set for the HistoryLoggerMock.Log")
		}
		return (*mm_results).err
	}
	if mmLog.funcLog != nil {
		return mmLog.funcLog(s)
	}
	mmLog.t.Fatalf("Unexpected call to HistoryLoggerMock.Log. %v", s)
	return
}

// LogAfterCounter returns a count of finished HistoryLoggerMock.Log invocations
func (mmLog *HistoryLoggerMock) LogAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLog.afterLogCounter)
}

// LogBeforeCounter returns a count of HistoryLoggerMock.Log invocations
func (mmLog *HistoryLoggerMock) LogBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLog.beforeLogCounter)
}

// Calls returns a list of arguments used in each call to HistoryLoggerMock.Log.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLog *mHistoryLoggerMockLog) Calls() []*HistoryLoggerMockLogParams {
	mmLog.mutex.RLock()

	argCopy := make([]*HistoryLoggerMockLogParams, len(mmLog.callArgs))
	copy(argCopy, mmLog.callArgs)

	mmLog.mutex.RUnlock()

	return argCopy
}

// MinimockLogDone returns true if the count of the Log invocations corresponds
// the number of defined expectations
func (m *HistoryLoggerMock) MinimockLogDone() bool {
	if m.LogMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.LogMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.LogMock.invocationsDone()
}

// MinimockLogInspect logs each unmet expectation
func (m *HistoryLoggerMock) MinimockLogInspect() {
	for _, e := range m.LogMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to HistoryLoggerMock.Log at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterLogCounter := mm_atomic.LoadUint64(&m.afterLogCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.LogMock.defaultExpectation != nil && afterLogCounter < 1 {
		if m.LogMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to HistoryLoggerMock.Log at\n%s", m.LogMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to HistoryLoggerMock.Log at\n%s with params: %#v", m.LogMock.defaultExpectation.expectationOrigins.origin, *m.LogMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLog != nil && afterLogCounter < 1 {
		m.t.Errorf("Expected call to HistoryLoggerMock.Log at\n%s", m.funcLogOrigin)
	}

	if !m.LogMock.invocationsDone() && afterLogCounter > 0 {
		m.t.Errorf("Expected %d calls to HistoryLoggerMock.Log at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.LogMock.expectedInvocations), m.LogMock.expectedInvocationsOrigin, afterLogCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *HistoryLoggerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockLogInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *HistoryLoggerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *HistoryLoggerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockLogDone()
}