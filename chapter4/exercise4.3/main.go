package main

import "sync"

type ReadWriteMutex struct {
	readerCounter int
	readersLock   sync.Mutex
	globalLock    sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
	// Only once entrance at a time
	rw.readersLock.Lock()
	rw.readerCounter++

	if rw.readerCounter == 1 {
		rw.globalLock.Lock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readerCounter--
	if rw.readerCounter == 0 {
		rw.globalLock.Unlock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) TryReadLock() bool {
	if rw.readersLock.TryLock() {
		globalBool := true
		if rw.readerCounter == 0 {
			globalBool = rw.globalLock.TryLock()
		}
		if globalBool {
			rw.readerCounter++
		}
		rw.readersLock.Unlock()
		return globalBool
	}
	return false
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}

func (rw *ReadWriteMutex) TryLock() bool {
	return rw.globalLock.TryLock()
}
