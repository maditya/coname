// Code generated by protoc-gen-gogo.
// source: config.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import github_com_maditya_protobuf_proto "github.com/maditya/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Config struct {
	Realms []*RealmConfig `protobuf:"bytes,1,rep,name=realms" json:"realms,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetRealms() []*RealmConfig {
	if m != nil {
		return m.Realms
	}
	return nil
}

type RealmConfig struct {
	// RealmName is the canonical name of the realm. It is signed by the
	// verifiers as a part of the epoch head.
	RealmName string `protobuf:"bytes,1,opt,name=RealmName,json=realmName,proto3" json:"RealmName,omitempty"`
	// Domains specifies a list of domains that belong to this realm.
	// Configuring one domain to belong to multiple realms is considered an
	// error.
	// TODO: support TLS-style wildcards.
	Domains []string `protobuf:"bytes,2,rep,name=domains" json:"domains,omitempty"`
	// Addr is the TCP (host:port) address of the keyserver GRPC interface.
	Addr string `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	// URL is the location of the secondary, HTTP-based interface to the
	// keyserver. It is not necessarily on the same host as addr.
	URL string `protobuf:"bytes,4,opt,name=URL,json=uRL,proto3" json:"URL,omitempty"`
	// VRFPublic is the public key of the verifiable random function used for
	// user id privacy. Here it is used to check that the anti-spam obfuscation
	// layer is properly used as a one-to-one mapping between real and
	// obfuscated usernames.
	VRFPublic []byte `protobuf:"bytes,5,opt,name=VRFPublic,json=vRFPublic,proto3" json:"VRFPublic,omitempty"`
	// VerificationPolicy specifies the conditions on how a lookup must be
	// verified for it to be accepted. Each verifier in VerificationPolicy MUST
	// have a NoOlderThan entry.
	VerificationPolicy *AuthorizationPolicy `protobuf:"bytes,6,opt,name=verification_policy,json=verificationPolicy" json:"verification_policy,omitempty"`
	// EpochTimeToLive specifies the duration for which an epoch is valid after
	// it has been issued. A client that has access to a clock MUST NOT accept
	// epoch heads with IssueTime more than EpochTimeToLive in the past.
	EpochTimeToLive Duration `protobuf:"bytes,7,opt,name=epoch_time_to_live,json=epochTimeToLive" json:"epoch_time_to_live"`
	// TreeNonce is the global nonce that is hashed into the Merkle tree nodes.
	TreeNonce []byte     `protobuf:"bytes,8,opt,name=tree_nonce,json=treeNonce,proto3" json:"tree_nonce,omitempty"`
	ClientTLS *TLSConfig `protobuf:"bytes,9,opt,name=client_tls,json=clientTls" json:"client_tls,omitempty"`
}

func (m *RealmConfig) Reset()                    { *m = RealmConfig{} }
func (*RealmConfig) ProtoMessage()               {}
func (*RealmConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *RealmConfig) GetVerificationPolicy() *AuthorizationPolicy {
	if m != nil {
		return m.VerificationPolicy
	}
	return nil
}

func (m *RealmConfig) GetEpochTimeToLive() Duration {
	if m != nil {
		return m.EpochTimeToLive
	}
	return Duration{}
}

func (m *RealmConfig) GetClientTLS() *TLSConfig {
	if m != nil {
		return m.ClientTLS
	}
	return nil
}

func init() {
	proto1.RegisterType((*Config)(nil), "proto.Config")
	proto1.RegisterType((*RealmConfig)(nil), "proto.RealmConfig")
}
func (this *Config) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Config)
	if !ok {
		that2, ok := that.(Config)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Config")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Config but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Config but is not nil && this == nil")
	}
	if len(this.Realms) != len(that1.Realms) {
		return fmt.Errorf("Realms this(%v) Not Equal that(%v)", len(this.Realms), len(that1.Realms))
	}
	for i := range this.Realms {
		if !this.Realms[i].Equal(that1.Realms[i]) {
			return fmt.Errorf("Realms this[%v](%v) Not Equal that[%v](%v)", i, this.Realms[i], i, that1.Realms[i])
		}
	}
	return nil
}
func (this *Config) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Config)
	if !ok {
		that2, ok := that.(Config)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Realms) != len(that1.Realms) {
		return false
	}
	for i := range this.Realms {
		if !this.Realms[i].Equal(that1.Realms[i]) {
			return false
		}
	}
	return true
}
func (this *RealmConfig) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*RealmConfig)
	if !ok {
		that2, ok := that.(RealmConfig)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *RealmConfig")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *RealmConfig but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *RealmConfig but is not nil && this == nil")
	}
	if this.RealmName != that1.RealmName {
		return fmt.Errorf("RealmName this(%v) Not Equal that(%v)", this.RealmName, that1.RealmName)
	}
	if len(this.Domains) != len(that1.Domains) {
		return fmt.Errorf("Domains this(%v) Not Equal that(%v)", len(this.Domains), len(that1.Domains))
	}
	for i := range this.Domains {
		if this.Domains[i] != that1.Domains[i] {
			return fmt.Errorf("Domains this[%v](%v) Not Equal that[%v](%v)", i, this.Domains[i], i, that1.Domains[i])
		}
	}
	if this.Addr != that1.Addr {
		return fmt.Errorf("Addr this(%v) Not Equal that(%v)", this.Addr, that1.Addr)
	}
	if this.URL != that1.URL {
		return fmt.Errorf("URL this(%v) Not Equal that(%v)", this.URL, that1.URL)
	}
	if !bytes.Equal(this.VRFPublic, that1.VRFPublic) {
		return fmt.Errorf("VRFPublic this(%v) Not Equal that(%v)", this.VRFPublic, that1.VRFPublic)
	}
	if !this.VerificationPolicy.Equal(that1.VerificationPolicy) {
		return fmt.Errorf("VerificationPolicy this(%v) Not Equal that(%v)", this.VerificationPolicy, that1.VerificationPolicy)
	}
	if !this.EpochTimeToLive.Equal(&that1.EpochTimeToLive) {
		return fmt.Errorf("EpochTimeToLive this(%v) Not Equal that(%v)", this.EpochTimeToLive, that1.EpochTimeToLive)
	}
	if !bytes.Equal(this.TreeNonce, that1.TreeNonce) {
		return fmt.Errorf("TreeNonce this(%v) Not Equal that(%v)", this.TreeNonce, that1.TreeNonce)
	}
	if !this.ClientTLS.Equal(that1.ClientTLS) {
		return fmt.Errorf("ClientTLS this(%v) Not Equal that(%v)", this.ClientTLS, that1.ClientTLS)
	}
	return nil
}
func (this *RealmConfig) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RealmConfig)
	if !ok {
		that2, ok := that.(RealmConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.RealmName != that1.RealmName {
		return false
	}
	if len(this.Domains) != len(that1.Domains) {
		return false
	}
	for i := range this.Domains {
		if this.Domains[i] != that1.Domains[i] {
			return false
		}
	}
	if this.Addr != that1.Addr {
		return false
	}
	if this.URL != that1.URL {
		return false
	}
	if !bytes.Equal(this.VRFPublic, that1.VRFPublic) {
		return false
	}
	if !this.VerificationPolicy.Equal(that1.VerificationPolicy) {
		return false
	}
	if !this.EpochTimeToLive.Equal(&that1.EpochTimeToLive) {
		return false
	}
	if !bytes.Equal(this.TreeNonce, that1.TreeNonce) {
		return false
	}
	if !this.ClientTLS.Equal(that1.ClientTLS) {
		return false
	}
	return true
}
func (this *Config) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&proto.Config{")
	if this.Realms != nil {
		s = append(s, "Realms: "+fmt.Sprintf("%#v", this.Realms)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RealmConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&proto.RealmConfig{")
	s = append(s, "RealmName: "+fmt.Sprintf("%#v", this.RealmName)+",\n")
	s = append(s, "Domains: "+fmt.Sprintf("%#v", this.Domains)+",\n")
	s = append(s, "Addr: "+fmt.Sprintf("%#v", this.Addr)+",\n")
	s = append(s, "URL: "+fmt.Sprintf("%#v", this.URL)+",\n")
	s = append(s, "VRFPublic: "+fmt.Sprintf("%#v", this.VRFPublic)+",\n")
	if this.VerificationPolicy != nil {
		s = append(s, "VerificationPolicy: "+fmt.Sprintf("%#v", this.VerificationPolicy)+",\n")
	}
	s = append(s, "EpochTimeToLive: "+strings.Replace(this.EpochTimeToLive.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "TreeNonce: "+fmt.Sprintf("%#v", this.TreeNonce)+",\n")
	if this.ClientTLS != nil {
		s = append(s, "ClientTLS: "+fmt.Sprintf("%#v", this.ClientTLS)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringConfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringConfig(e map[int32]github_com_maditya_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *Config) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Config) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Realms) > 0 {
		for _, msg := range m.Realms {
			data[i] = 0xa
			i++
			i = encodeVarintConfig(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *RealmConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RealmConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RealmName) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintConfig(data, i, uint64(len(m.RealmName)))
		i += copy(data[i:], m.RealmName)
	}
	if len(m.Domains) > 0 {
		for _, s := range m.Domains {
			data[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.Addr) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintConfig(data, i, uint64(len(m.Addr)))
		i += copy(data[i:], m.Addr)
	}
	if len(m.URL) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintConfig(data, i, uint64(len(m.URL)))
		i += copy(data[i:], m.URL)
	}
	if len(m.VRFPublic) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintConfig(data, i, uint64(len(m.VRFPublic)))
		i += copy(data[i:], m.VRFPublic)
	}
	if m.VerificationPolicy != nil {
		data[i] = 0x32
		i++
		i = encodeVarintConfig(data, i, uint64(m.VerificationPolicy.Size()))
		n1, err := m.VerificationPolicy.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	data[i] = 0x3a
	i++
	i = encodeVarintConfig(data, i, uint64(m.EpochTimeToLive.Size()))
	n2, err := m.EpochTimeToLive.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.TreeNonce) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintConfig(data, i, uint64(len(m.TreeNonce)))
		i += copy(data[i:], m.TreeNonce)
	}
	if m.ClientTLS != nil {
		data[i] = 0x4a
		i++
		i = encodeVarintConfig(data, i, uint64(m.ClientTLS.Size()))
		n3, err := m.ClientTLS.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeFixed64Config(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Config(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConfig(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedConfig(r randyConfig, easy bool) *Config {
	this := &Config{}
	if r.Intn(10) == 0 {
		v1 := r.Intn(5)
		this.Realms = make([]*RealmConfig, v1)
		for i := 0; i < v1; i++ {
			this.Realms[i] = NewPopulatedRealmConfig(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedRealmConfig(r randyConfig, easy bool) *RealmConfig {
	this := &RealmConfig{}
	this.RealmName = randStringConfig(r)
	v2 := r.Intn(10)
	this.Domains = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Domains[i] = randStringConfig(r)
	}
	this.Addr = randStringConfig(r)
	this.URL = randStringConfig(r)
	v3 := r.Intn(100)
	this.VRFPublic = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.VRFPublic[i] = byte(r.Intn(256))
	}
	if r.Intn(10) == 0 {
		this.VerificationPolicy = NewPopulatedAuthorizationPolicy(r, easy)
	}
	v4 := NewPopulatedDuration(r, easy)
	this.EpochTimeToLive = *v4
	v5 := r.Intn(100)
	this.TreeNonce = make([]byte, v5)
	for i := 0; i < v5; i++ {
		this.TreeNonce[i] = byte(r.Intn(256))
	}
	if r.Intn(10) != 0 {
		this.ClientTLS = NewPopulatedTLSConfig(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyConfig interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneConfig(r randyConfig) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringConfig(r randyConfig) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneConfig(r)
	}
	return string(tmps)
}
func randUnrecognizedConfig(r randyConfig, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldConfig(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldConfig(data []byte, r randyConfig, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateConfig(data, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		data = encodeVarintPopulateConfig(data, uint64(v7))
	case 1:
		data = encodeVarintPopulateConfig(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateConfig(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateConfig(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateConfig(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateConfig(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *Config) Size() (n int) {
	var l int
	_ = l
	if len(m.Realms) > 0 {
		for _, e := range m.Realms {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *RealmConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.RealmName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.Domains) > 0 {
		for _, s := range m.Domains {
			l = len(s)
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.VRFPublic)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.VerificationPolicy != nil {
		l = m.VerificationPolicy.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	l = m.EpochTimeToLive.Size()
	n += 1 + l + sovConfig(uint64(l))
	l = len(m.TreeNonce)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.ClientTLS != nil {
		l = m.ClientTLS.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Config) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Config{`,
		`Realms:` + strings.Replace(fmt.Sprintf("%v", this.Realms), "RealmConfig", "RealmConfig", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RealmConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RealmConfig{`,
		`RealmName:` + fmt.Sprintf("%v", this.RealmName) + `,`,
		`Domains:` + fmt.Sprintf("%v", this.Domains) + `,`,
		`Addr:` + fmt.Sprintf("%v", this.Addr) + `,`,
		`URL:` + fmt.Sprintf("%v", this.URL) + `,`,
		`VRFPublic:` + fmt.Sprintf("%v", this.VRFPublic) + `,`,
		`VerificationPolicy:` + strings.Replace(fmt.Sprintf("%v", this.VerificationPolicy), "AuthorizationPolicy", "AuthorizationPolicy", 1) + `,`,
		`EpochTimeToLive:` + strings.Replace(strings.Replace(this.EpochTimeToLive.String(), "Duration", "Duration", 1), `&`, ``, 1) + `,`,
		`TreeNonce:` + fmt.Sprintf("%v", this.TreeNonce) + `,`,
		`ClientTLS:` + strings.Replace(fmt.Sprintf("%v", this.ClientTLS), "TLSConfig", "TLSConfig", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Config) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Realms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Realms = append(m.Realms, &RealmConfig{})
			if err := m.Realms[len(m.Realms)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RealmConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RealmConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RealmConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RealmName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RealmName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domains = append(m.Domains, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VRFPublic", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VRFPublic = append(m.VRFPublic[:0], data[iNdEx:postIndex]...)
			if m.VRFPublic == nil {
				m.VRFPublic = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VerificationPolicy == nil {
				m.VerificationPolicy = &AuthorizationPolicy{}
			}
			if err := m.VerificationPolicy.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochTimeToLive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EpochTimeToLive.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreeNonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TreeNonce = append(m.TreeNonce[:0], data[iNdEx:postIndex]...)
			if m.TreeNonce == nil {
				m.TreeNonce = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientTLS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClientTLS == nil {
				m.ClientTLS = &TLSConfig{}
			}
			if err := m.ClientTLS.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x51, 0xbf, 0x6f, 0x13, 0x31,
	0x18, 0x4d, 0x9a, 0x34, 0xc5, 0x4e, 0x20, 0x95, 0x59, 0xac, 0x08, 0x52, 0xd4, 0xa9, 0x62, 0x48,
	0x50, 0x61, 0x60, 0x42, 0xe2, 0x8a, 0x58, 0x38, 0x55, 0x95, 0x39, 0x58, 0x4f, 0xf7, 0xc3, 0xb9,
	0x58, 0xba, 0x3b, 0x47, 0x3e, 0x5f, 0xa4, 0x32, 0xf1, 0xe7, 0xf0, 0x27, 0x30, 0x32, 0x76, 0xec,
	0xc8, 0x54, 0xb5, 0x99, 0x18, 0x33, 0xc2, 0xc6, 0x77, 0x9f, 0x2f, 0x52, 0x86, 0x27, 0xfb, 0xbd,
	0x7b, 0xef, 0xf9, 0xec, 0x8f, 0x8e, 0x12, 0x5d, 0x2e, 0x54, 0x36, 0x5b, 0x19, 0x6d, 0x35, 0x3b,
	0xc4, 0x65, 0xf2, 0x2a, 0x53, 0x76, 0x59, 0xc7, 0xb3, 0x44, 0x17, 0xf3, 0x22, 0x4a, 0x95, 0xbd,
	0x8e, 0xe6, 0xf8, 0x25, 0xae, 0x17, 0xf3, 0x4c, 0x67, 0x1a, 0x09, 0xee, 0x5c, 0x70, 0x32, 0x4a,
	0x72, 0x25, 0x4b, 0xdb, 0xb2, 0x27, 0x69, 0x6d, 0x22, 0xab, 0x74, 0xd9, 0xf2, 0xb1, 0xcd, 0xab,
	0xfd, 0x73, 0x4e, 0xdf, 0xd0, 0xc1, 0x05, 0x72, 0xf6, 0x92, 0x0e, 0x8c, 0x8c, 0xf2, 0xa2, 0xe2,
	0xdd, 0x17, 0xbd, 0xb3, 0xe1, 0x39, 0x73, 0x8e, 0x99, 0x68, 0x44, 0xe7, 0x11, 0xad, 0xe3, 0xf4,
	0xdf, 0x01, 0x1d, 0xee, 0xe9, 0xec, 0x19, 0x25, 0x48, 0x2f, 0xa3, 0x42, 0x42, 0xbc, 0x7b, 0x46,
	0x04, 0x31, 0x3b, 0x81, 0x71, 0x7a, 0x94, 0xea, 0x22, 0x52, 0x65, 0xc5, 0x0f, 0xa0, 0x9a, 0x88,
	0x1d, 0x65, 0x8c, 0xf6, 0xa3, 0x34, 0x35, 0xbc, 0x87, 0x11, 0xdc, 0xb3, 0x63, 0xda, 0xfb, 0x22,
	0x7c, 0xde, 0x47, 0xa9, 0x57, 0x0b, 0xbf, 0x69, 0xff, 0x2a, 0x3e, 0x5e, 0xd5, 0x71, 0xae, 0x12,
	0x7e, 0x08, 0xfa, 0x48, 0x90, 0xf5, 0x4e, 0x60, 0x9f, 0xe8, 0xd3, 0xb5, 0x34, 0x6a, 0xa1, 0x12,
	0xbc, 0x68, 0xb8, 0xd2, 0xa0, 0x5e, 0xf3, 0x01, 0xf8, 0x86, 0xe7, 0x93, 0xf6, 0x12, 0xef, 0x6b,
	0xbb, 0xd4, 0x46, 0x7d, 0x43, 0xcb, 0x15, 0x3a, 0x04, 0xdb, 0x8f, 0x39, 0x8d, 0x79, 0x94, 0xc9,
	0x95, 0x4e, 0x96, 0xa1, 0x55, 0x85, 0x0c, 0xad, 0x0e, 0x73, 0xb5, 0x96, 0xfc, 0x08, 0xbb, 0xc6,
	0x6d, 0xd7, 0x87, 0xf6, 0x49, 0xbd, 0xfe, 0xcd, 0xdd, 0x49, 0x47, 0x8c, 0x31, 0x10, 0x80, 0x3f,
	0xd0, 0x3e, 0xb8, 0xd9, 0x73, 0x4a, 0xad, 0x91, 0x32, 0x2c, 0x75, 0x99, 0x48, 0xfe, 0xc8, 0xfd,
	0x6f, 0xa3, 0x5c, 0x36, 0x02, 0x7b, 0x47, 0xa9, 0x1b, 0x51, 0x08, 0xb3, 0xe0, 0x04, 0xab, 0x8f,
	0xdb, 0xea, 0xc0, 0xff, 0xec, 0x5e, 0xd4, 0x7b, 0xbc, 0xb9, 0x3b, 0x21, 0x17, 0xe8, 0x03, 0x51,
	0x10, 0x17, 0x09, 0xf2, 0xca, 0x7b, 0x7b, 0xfb, 0x30, 0xed, 0xfc, 0x06, 0xdc, 0x3f, 0x4c, 0xbb,
	0x5b, 0xc0, 0x5f, 0xc0, 0xf7, 0xcd, 0xb4, 0xfb, 0x03, 0xf0, 0x13, 0xf0, 0x0b, 0x70, 0x03, 0xb8,
	0x05, 0xdc, 0x03, 0xfe, 0x6c, 0xa6, 0x9d, 0x2d, 0xac, 0xf1, 0x00, 0x0f, 0x79, 0xfd, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0xdc, 0xed, 0x4e, 0x79, 0x6a, 0x02, 0x00, 0x00,
}
