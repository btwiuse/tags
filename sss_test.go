package tags

import (
	"flag"
	"reflect"
	"testing"
)

func TestSpaceSeparatedStrings_String(t *testing.T) {
	tests := []struct {
		name string
		tags SpaceSeparatedStrings
		want string
	}{
		{
			name: "empty slice",
			tags: SpaceSeparatedStrings{},
			want: "",
		},
		{
			name: "single element",
			tags: SpaceSeparatedStrings{"hello"},
			want: "hello",
		},
		{
			name: "multiple elements",
			tags: SpaceSeparatedStrings{"hello", "world"},
			want: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tags.String(); got != tt.want {
				t.Errorf("SpaceSeparatedStrings.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpaceSeparatedStrings_Set(t *testing.T) {
	tests := []struct {
		name    string
		tags    *SpaceSeparatedStrings
		value   string
		wantErr bool
		want    SpaceSeparatedStrings
	}{
		{
			name:    "empty string",
			tags:    &SpaceSeparatedStrings{},
			value:   "",
			wantErr: false,
			want:    SpaceSeparatedStrings{},
		},
		{
			name:    "single word",
			tags:    &SpaceSeparatedStrings{},
			value:   "hello",
			wantErr: false,
			want:    SpaceSeparatedStrings{"hello"},
		},
		{
			name:    "multiple words",
			tags:    &SpaceSeparatedStrings{},
			value:   "hello world",
			wantErr: false,
			want:    SpaceSeparatedStrings{"hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tags.Set(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("SpaceSeparatedStrings.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(*tt.tags, tt.want) {
				t.Errorf("SpaceSeparatedStrings after Set() = %v, want %v", *tt.tags, tt.want)
			}
		})
	}
}

func TestSpaceSeparatedStrings_Flag(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    SpaceSeparatedStrings
		wantErr bool
	}{
		{
			name:    "no args",
			args:    []string{},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "single arg",
			args:    []string{"-tags=hello"},
			want:    SpaceSeparatedStrings{"hello"},
			wantErr: false,
		},
		{
			name:    "multiple args",
			args:    []string{"-tags=hello", "-tags=world"},
			want:    SpaceSeparatedStrings{"hello", "world"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tags SpaceSeparatedStrings
			flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
			flagSet.Var(&tags, "tags", "Space separated tags")

			err := flagSet.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Error parsing flags = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tags, tt.want) {
				t.Errorf("SpaceSeparatedStrings after parsing flags = %v, want %v", tags, tt.want)
			}
		})
	}
}
