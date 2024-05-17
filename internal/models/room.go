package models

import (
	"crypto/rand"
	"sync"
)

type Room struct {
	// ID is the unique identifier of the room
	ID string

	// BoardID is the unique identifier of the board that the room belongs to
	BoardID string

	// Users is a map of users in the room
	Users []*User

	// LeaderID is the unique identifier of the leader of the room
	LeaderID string

	// Elements is a string that represents the elements of the board
	Elements string

	// AppState is a string that represents the app state of the board
	AppState string

	// mtx is a mutex
	mtx *sync.RWMutex

	RoomMutex *sync.Mutex
}

// NewRoom creates a new room.
func NewRoom(boardID string) *Room {
	return &Room{
		ID:        generateRandomID(),
		BoardID:   boardID,
		Users:     make([]*User, 0),
		LeaderID:  "0",
		mtx:       &sync.RWMutex{},
		RoomMutex: &sync.Mutex{},
	}
}

func (r *Room) AddUser(newUser *User) {
	// Add user to the room
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.Users = append(r.Users, newUser)
}

func (r *Room) RemoveUser(userID string) {
	// Remove user from the room
	r.mtx.Lock()
	defer r.mtx.Unlock()
	for i, u := range r.Users {
		if u.ID == userID {
			r.Users = append(r.Users[:i], r.Users[i+1:]...)
			break
		}
	}
}

func (r *Room) GetUsers() []*User {
	// Get users of the room
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.Users
}

func (r *Room) SetLeader(leaderID string) {
	// Set leader of the room
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.LeaderID = leaderID
}

func (r *Room) SetElements(elements string) {
	// Set elements of the room
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.Elements = elements
}

func (r *Room) SetAppState(appState string) {
	// Set app state of the room
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.AppState = appState
}

func (r *Room) GetElements() string {
	// Get elements of the room
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.Elements
}

func (r *Room) GetAppState() string {
	// Get app state of the room
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.AppState
}

// generateRandomID generates a random ID for the room.
func generateRandomID() string {
	const idLength = 16
	const idChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, idLength)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	for i := 0; i < idLength; i++ {
		b[i] = idChars[int(b[i])%len(idChars)]
	}

	return string(b)
}
