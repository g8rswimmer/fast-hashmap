package hashmap

import (
	"reflect"
	"testing"
)

func TestHashMap_Get(t *testing.T) {
	type fields struct {
		table []entries
	}
	type args struct {
		k int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "One Entry",
			fields: fields{
				table: []entries{
					nil,
					entries{
						entry{key: 1, obj: "hello"},
					},
					nil,
				},
			},
			args: args{
				k: 1,
			},
			want:    "hello",
			wantErr: false,
		},
		{
			name: "Two Entry",
			fields: fields{
				table: []entries{
					nil,
					entries{
						entry{key: 1, obj: "hello"},
						entry{key: 1025, obj: "good-bye"},
					},
					nil,
				},
			},
			args: args{
				k: 1025,
			},
			want:    "good-bye",
			wantErr: false,
		},
		{
			name: "not found",
			fields: fields{
				table: []entries{
					nil,
					entries{
						entry{key: 1, obj: "hello"},
						entry{key: 1025, obj: "good-bye"},
					},
					nil,
				},
			},
			args: args{
				k: 1024,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not found too",
			fields: fields{
				table: []entries{
					nil,
					entries{
						entry{key: 1, obj: "hello"},
						entry{key: 1025, obj: "good-bye"},
					},
					nil,
				},
			},
			args: args{
				k: 2049,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				table: tt.fields.table,
			}
			got, err := h.Get(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashMap.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testHashMap() *HashMap {
	hm := New()
	hm.Put(2, "what.")
	hm.Put(12, "what..")
	hm.Put(112, "what...")
	hm.Put(1112, "what....")
	hm.Put(11112, "what.....")
	hm.Put(111112, "what........")
	return hm
}
func TestHashMap_Put(t *testing.T) {
	type args struct {
		k int
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Put Empty",
			args: args{
				k: 3,
				v: "yep",
			},
			want: "yep",
		},
		{
			name: "Put with entry",
			args: args{
				k: 1026,
				v: "yep",
			},
			want: "yep",
		},
		{
			name: "Put oever",
			args: args{
				k: 2,
				v: "yep",
			},
			want: "yep",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := testHashMap()
			if err := h.Put(tt.args.k, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("HashMap.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err := h.Get(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashMap.Put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashMap.Put() = %v, want %v", got, tt.want)
			}
		})
	}
}
