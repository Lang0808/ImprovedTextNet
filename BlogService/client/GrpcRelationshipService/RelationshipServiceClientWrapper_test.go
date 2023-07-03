package GrpcRelationshipService

import "testing"

import "github.com/magiconair/properties/assert"

func TestGetRelationshipFromNumber(t *testing.T) {
	tests := []struct {
		desc     string
		n        int32
		expected RelationshipType
	}{
		{"NONE", 0, RelationshipType_NONE},
		{"FRIEND", 1, RelationshipType_FRIEND},
		{"FOLLOWING", 2, RelationshipType_FOLLOWING},
		{"FOLLOWED", 3, RelationshipType_FOLLOWED},
		{"BLOCKING", 4, RelationshipType_BLOCKING},
		{"BLOCKED", 5, RelationshipType_BLOCKED},
		{"SENT_REQUEST", 6, RelationshipType_SENT_REQUEST},
		{"RECEIVED_REQUEST", 7, RelationshipType_RECEIVED_REQUEST},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := GetRelationshipFromNumber(test.n)
			if err != nil {
				t.Fatalf("%v; Error: %v\n", test.desc, err)
			}
			assert.Equal(t, test.expected, *got, "Incorrect value")
		})
	}
}
