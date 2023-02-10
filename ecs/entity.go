package ecs

// Entity identifier.
// Holds an entity ID and it's generation for recycling.
//
// Entities should only be created via the [World], using [World.NewEntity].
type Entity struct {
	id  eid
	gen uint16
}

func newEntity(id eid) Entity {
	return Entity{id, 0}
}

// IsZero returns whether this entity is the reserved zero entity.
func (e Entity) IsZero() bool {
	return e.id == 0
}

type entityIndex struct {
	arch  *archetype
	index uint32
}
