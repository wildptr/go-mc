package bot

import (
	"github.com/google/uuid"
	"github.com/wildptr/go-mc/bot/world/entity"
	"github.com/wildptr/go-mc/chat"

	pk "github.com/wildptr/go-mc/net/packet"
)

type eventBroker struct {
	GameStart      func() error
	ChatMsg        func(msg chat.Message, pos byte) error
	Disconnect     func(reason chat.Message) error
	HealthChange   func() error
	Die            func() error
	SoundPlay      func(name string, category int, x, y, z float64, volume, pitch float32) error
	PluginMessage  func(channel string, data []byte) error
	HeldItemChange func(slot int) error

	WindowsItem       func(id byte, slots []entity.Slot) error
	WindowsItemChange func(id byte, slotID int, slot entity.Slot) error

	SpawnObject func(entity_id int, obj_uuid uuid.UUID, typ int, x, y, z float64, pitch, yaw int8, data int, vx, vy, vz int) error

	// ReceivePacket will be called when new packet arrive.
	// Default handler will run only if pass == false.
	ReceivePacket func(p pk.Packet) (pass bool, err error)
}
