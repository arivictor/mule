package mule

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_mule_Unique(t *testing.T) {
	type fields struct {
		errors map[string]string
	}
	type args struct {
		values []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "no values should pass",
			fields: fields{errors: make(map[string]string)},
			args:   args{values: []string{}},
			want:   true,
		},
		{
			name:   "unique values should pass",
			fields: fields{errors: make(map[string]string)},
			args:   args{values: []string{"a", "b", "c"}},
			want:   true,
		},
		{
			name:   "non-unique values should fail",
			fields: fields{errors: make(map[string]string)},
			args:   args{values: []string{"a", "b", "a"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mule{
				errors: tt.fields.errors,
			}
			if got := m.Unique(tt.args.values); got != tt.want {
				t.Errorf("mule.Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mule_Matches(t *testing.T) {
	emailRegEx := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	type fields struct {
		errors map[string]string
	}
	type args struct {
		value string
		rx    *regexp.Regexp
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "matching values should pass",
			fields: fields{errors: make(map[string]string)},
			args:   args{value: "apple", rx: regexp.MustCompile("apple")},
			want:   true,
		},
		{
			name:   "non-matching values should fail",
			fields: fields{errors: make(map[string]string)},
			args:   args{value: "banana", rx: regexp.MustCompile("apple")},
			want:   false,
		},
		{
			name:   "matching values with regex should pass",
			fields: fields{errors: make(map[string]string)},
			args:   args{value: "me@test.com", rx: emailRegEx},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mule{
				errors: tt.fields.errors,
			}
			if got := m.Matches(tt.args.value, tt.args.rx); got != tt.want {
				t.Errorf("mule.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mule_In(t *testing.T) {
	type fields struct {
		errors map[string]string
	}
	type args struct {
		value string
		list  []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "value in list should pass",
			fields: fields{errors: make(map[string]string)},
			args:   args{value: "a", list: []string{"a", "b", "c"}},
			want:   true,
		},
		{
			name:   "value not in list should fail",
			fields: fields{errors: make(map[string]string)},
			args:   args{value: "a", list: []string{"x", "b", "c"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mule{
				errors: tt.fields.errors,
			}
			if got := m.In(tt.args.value, tt.args.list...); got != tt.want {
				t.Errorf("mule.In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mule_Check(t *testing.T) {
	type fields struct {
		errors map[string]string
	}
	type args struct {
		ok      bool
		key     string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "empty map and should pass",
			fields: fields{errors: make(map[string]string)},
			args:   args{ok: true, key: "test", message: "testing"},
		},
		{
			name:   "overwrites existing key but should pass",
			fields: fields{errors: map[string]string{"test": "testing"}},
			args:   args{ok: true, key: "test", message: "testing"},
		},
		{
			name:   "creates new key and should pass",
			fields: fields{errors: map[string]string{"test": "testing"}},
			args:   args{ok: true, key: "test1", message: "testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mule{
				errors: tt.fields.errors,
			}
			m.Check(tt.args.ok, tt.args.key, tt.args.message)
		})
	}
}

func Test_mule_addError(t *testing.T) {
	type fields struct {
		errors map[string]string
	}
	type args struct {
		key     string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "should pass",
			fields: fields{errors: make(map[string]string)},
			args:   args{key: "hello", message: "world"},
		},
		{
			name:   "should pass",
			fields: fields{errors: make(map[string]string)},
			args:   args{key: "", message: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mule{
				errors: tt.fields.errors,
			}
			m.addError(tt.args.key, tt.args.message)
		})
	}
}

func Test_mule_Valid(t *testing.T) {
	type fields struct {
		errors map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "no errors should pass",
			fields: fields{errors: make(map[string]string)},
			want:   true,
		},
		{
			name:   "with errors should fail",
			fields: fields{errors: map[string]string{"test": "test"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &mule{
				errors: tt.fields.errors,
			}
			if got := v.Valid(); got != tt.want {
				t.Errorf("mule.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mule_Errors(t *testing.T) {
	type fields struct {
		errors map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name:   "should return empty",
			fields: fields{errors: make(map[string]string)},
			want:   make(map[string]string),
		},
		{
			name:   "should return not empty",
			fields: fields{errors: map[string]string{"test": "test"}},
			want:   map[string]string{"test": "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mule{
				errors: tt.fields.errors,
			}
			if got := m.Errors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mule.Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Mule
	}{
		{
			name: "should return new mul with new errors map",
			want: &mule{errors: make(map[string]string)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
