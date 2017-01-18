// Code generated by protoc-gen-gogo.
// source: server/pkg/hashtree/hashtree.proto
// DO NOT EDIT!

/*
Package hashtree is a generated protocol buffer package.

It is generated from these files:
	server/pkg/hashtree/hashtree.proto

It has these top-level messages:
	FileNode
	DirectoryNode
	Node
	HashTree
*/
package hashtree

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A node corresponding to a file (which is also a leaf node).
type FileNode struct {
	// Block refs pointing to the file's contents in block storage
	BlockRefs []*pfs.BlockRef `protobuf:"bytes,3,rep,name=block_refs,json=blockRefs" json:"block_refs,omitempty"`
}

func (m *FileNode) Reset()                    { *m = FileNode{} }
func (m *FileNode) String() string            { return proto.CompactTextString(m) }
func (*FileNode) ProtoMessage()               {}
func (*FileNode) Descriptor() ([]byte, []int) { return fileDescriptorHashtree, []int{0} }

func (m *FileNode) GetBlockRefs() []*pfs.BlockRef {
	if m != nil {
		return m.BlockRefs
	}
	return nil
}

// A node corresponding to a directory.
type DirectoryNode struct {
	// Children of this directory. Note that paths are relative, so if "/foo/bar"
	// has a child "baz", that means that there is a file at "/foo/bar/baz"
	Children []string `protobuf:"bytes,3,rep,name=children" json:"children,omitempty"`
}

func (m *DirectoryNode) Reset()                    { *m = DirectoryNode{} }
func (m *DirectoryNode) String() string            { return proto.CompactTextString(m) }
func (*DirectoryNode) ProtoMessage()               {}
func (*DirectoryNode) Descriptor() ([]byte, []int) { return fileDescriptorHashtree, []int{1} }

func (m *DirectoryNode) GetChildren() []string {
	if m != nil {
		return m.Children
	}
	return nil
}

// A node (either a file or a directory)
type Node struct {
	// Name of the file/directory.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Hash of the name and contents. This can be used to detect if the name
	// or contents have changed between versions.
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// Size of the node; if this is a directory, the size includes all children
	Size_ uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// Exactly one of the following fields must be set. The type of this node will
	// be determined by which field is set.
	FileNode *FileNode      `protobuf:"bytes,4,opt,name=file_node,json=fileNode" json:"file_node,omitempty"`
	DirNode  *DirectoryNode `protobuf:"bytes,5,opt,name=dir_node,json=dirNode" json:"dir_node,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorHashtree, []int{2} }

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Node) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Node) GetFileNode() *FileNode {
	if m != nil {
		return m.FileNode
	}
	return nil
}

func (m *Node) GetDirNode() *DirectoryNode {
	if m != nil {
		return m.DirNode
	}
	return nil
}

// A Tree corresponding to the complete contents of a pachyderm repo at a given
// commit (based on a Merkle Tree). We store one HashTree for every PFS commit.
type HashTree struct {
	// Arbitrary version number, set by the corresponding library in hashtree.go.
	// This ensures that if the hash function used to create these trees is
	// changed, we won't run into errors when deserializing old trees.
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Map from each node's path to the Node message with that node's details.
	// See "Potential Optimizations" at the end for a compression scheme that
	// could be useful if this map gets too large.
	//
	// Note that the key must end in "/" if an only if the value has .dir_node set
	// (i.e. iff the path points to a directory).
	Fs map[string]*Node `protobuf:"bytes,2,rep,name=fs" json:"fs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HashTree) Reset()                    { *m = HashTree{} }
func (m *HashTree) String() string            { return proto.CompactTextString(m) }
func (*HashTree) ProtoMessage()               {}
func (*HashTree) Descriptor() ([]byte, []int) { return fileDescriptorHashtree, []int{3} }

func (m *HashTree) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HashTree) GetFs() map[string]*Node {
	if m != nil {
		return m.Fs
	}
	return nil
}

func init() {
	proto.RegisterType((*FileNode)(nil), "FileNode")
	proto.RegisterType((*DirectoryNode)(nil), "DirectoryNode")
	proto.RegisterType((*Node)(nil), "Node")
	proto.RegisterType((*HashTree)(nil), "HashTree")
}

func init() { proto.RegisterFile("server/pkg/hashtree/hashtree.proto", fileDescriptorHashtree) }

var fileDescriptorHashtree = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x6b, 0xe3, 0x30,
	0x14, 0xc4, 0x91, 0xff, 0x6c, 0xec, 0x97, 0xcd, 0xb2, 0x2b, 0xf6, 0x20, 0xd2, 0x8b, 0xeb, 0x43,
	0x71, 0x69, 0x51, 0x20, 0xbd, 0x84, 0xd2, 0x53, 0x69, 0x43, 0x4f, 0x3d, 0x88, 0xde, 0x83, 0x63,
	0x3f, 0xd5, 0x22, 0xae, 0x15, 0x24, 0x37, 0x90, 0xde, 0xfb, 0x11, 0xfa, 0x7d, 0x8b, 0x94, 0x38,
	0xd0, 0x83, 0xe1, 0xf7, 0x86, 0x19, 0x98, 0xb1, 0x20, 0xb7, 0x68, 0x76, 0x68, 0x66, 0xdb, 0xcd,
	0xeb, 0xac, 0x29, 0x6d, 0xd3, 0x1b, 0xc4, 0x13, 0xf0, 0xad, 0xd1, 0xbd, 0x9e, 0xfe, 0xaf, 0x5a,
	0x85, 0x5d, 0x3f, 0xdb, 0x4a, 0xeb, 0xbe, 0x83, 0x9a, 0x2f, 0x20, 0x59, 0xaa, 0x16, 0x9f, 0x75,
	0x8d, 0xf4, 0x1a, 0x60, 0xdd, 0xea, 0x6a, 0xb3, 0x32, 0x28, 0x2d, 0x0b, 0xb3, 0xb0, 0x18, 0xcf,
	0x27, 0xdc, 0x79, 0xef, 0x9d, 0x2c, 0x50, 0x8a, 0x74, 0x7d, 0x24, 0x9b, 0x5f, 0xc1, 0xe4, 0x41,
	0x19, 0xac, 0x7a, 0x6d, 0xf6, 0x3e, 0x3e, 0x85, 0xa4, 0x6a, 0x54, 0x5b, 0x1b, 0xec, 0x7c, 0x38,
	0x15, 0xa7, 0x3b, 0xff, 0x22, 0x10, 0x79, 0x13, 0x85, 0xa8, 0x2b, 0xdf, 0x90, 0x91, 0x8c, 0x14,
	0xa9, 0xf0, 0xec, 0x34, 0xd7, 0x95, 0x05, 0x19, 0x29, 0x7e, 0x0b, 0xcf, 0x4e, 0xb3, 0xea, 0x03,
	0x59, 0x98, 0x91, 0x22, 0x12, 0x9e, 0xe9, 0x05, 0xa4, 0x52, 0xb5, 0xb8, 0xea, 0x74, 0x8d, 0x2c,
	0xca, 0x48, 0x31, 0x9e, 0xa7, 0x7c, 0x68, 0x2f, 0x12, 0x39, 0xec, 0xb8, 0x84, 0xa4, 0x56, 0xe6,
	0x60, 0x8b, 0xbd, 0xed, 0x0f, 0xff, 0x51, 0x55, 0x8c, 0x6a, 0x65, 0x1c, 0xe4, 0x9f, 0x04, 0x92,
	0xa7, 0xd2, 0x36, 0x2f, 0x06, 0x91, 0x32, 0x18, 0xed, 0xd0, 0x58, 0xa5, 0x3b, 0x5f, 0x2f, 0x16,
	0xc3, 0x49, 0xcf, 0x21, 0x90, 0x96, 0x05, 0xfe, 0x8f, 0xfc, 0xe3, 0x43, 0x80, 0x2f, 0xed, 0x63,
	0xd7, 0x9b, 0xbd, 0x08, 0xa4, 0x9d, 0xde, 0xc1, 0xe8, 0x78, 0xd2, 0xbf, 0x10, 0x6e, 0x70, 0x7f,
	0x9c, 0xe8, 0x90, 0x9e, 0x41, 0xbc, 0x2b, 0xdb, 0x77, 0xf4, 0x13, 0xc7, 0xf3, 0x98, 0xfb, 0x16,
	0x07, 0xed, 0x36, 0x58, 0x90, 0xf5, 0x2f, 0xff, 0x1a, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xa1, 0x47, 0x98, 0x4f, 0xc9, 0x01, 0x00, 0x00,
}