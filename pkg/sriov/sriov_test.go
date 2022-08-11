package sriov

import (
	"net"

	sriovtypes "github.com/MangoBoost/sriov-cni/pkg/types"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/containernetworking/plugins/pkg/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/vishvananda/netlink"
)

// Code generated by mockery v1.0.0. DO NOT EDIT.
// MockNetlinkManager is an autogenerated mock type for the NetlinkManager type
type MockNetlinkManager struct {
	mock.Mock
}

// LinkByName provides a mock function with given fields: _a0
func (_m *MockNetlinkManager) LinkByName(_a0 string) (netlink.Link, error) {
	ret := _m.Called(_a0)

	var r0 netlink.Link
	if rf, ok := ret.Get(0).(func(string) netlink.Link); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(netlink.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkSetVfVlan provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockNetlinkManager) LinkSetVfVlan(_a0 netlink.Link, _a1 int, _a2 int) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, int) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfHardwareAddr provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockNetlinkManager) LinkSetVfHardwareAddr(_a0 netlink.Link, _a1 int, _a2 net.HardwareAddr) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, net.HardwareAddr) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetHardwareAddr provides a mock function with given fields: _a0, _a1
func (_m *MockNetlinkManager) LinkSetHardwareAddr(_a0 netlink.Link, _a1 net.HardwareAddr) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, net.HardwareAddr) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfVlanQos provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockNetlinkManager) LinkSetVfVlanQos(_a0 netlink.Link, _a1 int, _a2 int, _a3 int) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, int, int) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetUp provides a mock function with given fields: _a0
func (_m *MockNetlinkManager) LinkSetUp(_a0 netlink.Link) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetDown provides a mock function with given fields: _a0
func (_m *MockNetlinkManager) LinkSetDown(_a0 netlink.Link) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetNsFd provides a mock function with given fields: _a0, _a1
func (_m *MockNetlinkManager) LinkSetNsFd(_a0 netlink.Link, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetName provides a mock function with given fields: _a0, _a1
func (_m *MockNetlinkManager) LinkSetName(_a0 netlink.Link, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfRate provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockNetlinkManager) LinkSetVfRate(_a0 netlink.Link, _a1 int, _a2 int, _a3 int) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, int, int) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfSpoofchk provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockNetlinkManager) LinkSetVfSpoofchk(_a0 netlink.Link, _a1 int, _a2 bool) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, bool) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfTrust provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockNetlinkManager) LinkSetVfTrust(_a0 netlink.Link, _a1 int, _a2 bool) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, bool) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfState provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockNetlinkManager) LinkSetVfState(_a0 netlink.Link, _a1 int, _a2 uint32) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, uint32) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FakeLink is a dummy netlink struct used during testing
type FakeLink struct {
	netlink.LinkAttrs
}

// type FakeLink struct {
// 	linkAtrrs *netlink.LinkAttrs
// }

func (l *FakeLink) Attrs() *netlink.LinkAttrs {
	return &l.LinkAttrs
}

func (l *FakeLink) Type() string {
	return "FakeLink"
}

var _ = Describe("Sriov", func() {
	var (
		t GinkgoTInterface
	)
	BeforeEach(func() {
		t = GinkgoT()
	})

	Context("Checking SetupVF function", func() {
		var (
			podifName string
			contID    string
			netconf   *sriovtypes.NetConf
		)

		BeforeEach(func() {
			podifName = "net1"
			contID = "dummycid"
			netconf = &sriovtypes.NetConf{
				Master:      "enp175s0f1",
				DeviceID:    "0000:af:06.0",
				VFID:        0,
				ContIFNames: "net1",
				OrigVfState: sriovtypes.VfState{
					HostIFName: "enp175s6",
				},
			}
			t = GinkgoT()
		})

		It("Assuming existing interface", func() {
			var targetNetNS ns.NetNS
			targetNetNS, err := testutils.NewNS()
			defer func() {
				if targetNetNS != nil {
					targetNetNS.Close()
				}
			}()
			Expect(err).NotTo(HaveOccurred())
			mocked := &MockNetlinkManager{}
			fakeMac, err := net.ParseMAC("6e:16:06:0e:b7:e9")

			Expect(err).NotTo(HaveOccurred())

			fakeLink := &FakeLink{netlink.LinkAttrs{
				Index:        1000,
				Name:         "dummylink",
				HardwareAddr: fakeMac,
			}}

			mocked.On("LinkByName", mock.AnythingOfType("string")).Return(fakeLink, nil)
			mocked.On("LinkSetDown", fakeLink).Return(nil)
			mocked.On("LinkSetName", fakeLink, mock.Anything).Return(nil)
			mocked.On("LinkSetNsFd", fakeLink, mock.AnythingOfType("int")).Return(nil)
			mocked.On("LinkSetUp", fakeLink).Return(nil)
			mocked.On("LinkSetVfVlan", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
			mocked.On("LinkSetVfVlanQos", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
			sm := sriovManager{nLink: mocked}
			macAddr, err := sm.SetupVF(netconf, podifName, contID, targetNetNS)
			Expect(err).NotTo(HaveOccurred())
			Expect(macAddr).To(Equal("6e:16:06:0e:b7:e9"))
		})
		It("Setting VF's MAC address", func() {
			var targetNetNS ns.NetNS
			targetNetNS, err := testutils.NewNS()
			defer func() {
				if targetNetNS != nil {
					targetNetNS.Close()
				}
			}()
			Expect(err).NotTo(HaveOccurred())
			mocked := &MockNetlinkManager{}
			fakeMac, err := net.ParseMAC("6e:16:06:0e:b7:e9")
			Expect(err).NotTo(HaveOccurred())

			netconf.MAC = "e4:11:22:33:44:55"
			expMac, err := net.ParseMAC(netconf.MAC)
			Expect(err).NotTo(HaveOccurred())

			fakeLink := &FakeLink{netlink.LinkAttrs{
				Index:        1000,
				Name:         "dummylink",
				HardwareAddr: fakeMac,
			}}

			mocked.On("LinkByName", mock.AnythingOfType("string")).Return(fakeLink, nil)
			mocked.On("LinkSetDown", fakeLink).Return(nil)
			mocked.On("LinkSetName", fakeLink, mock.Anything).Return(nil)
			mocked.On("LinkSetHardwareAddr", fakeLink, expMac).Return(nil)
			mocked.On("LinkSetNsFd", fakeLink, mock.AnythingOfType("int")).Return(nil)
			mocked.On("LinkSetUp", fakeLink).Return(nil)
			sm := sriovManager{nLink: mocked}
			macAddr, err := sm.SetupVF(netconf, podifName, contID, targetNetNS)
			Expect(err).NotTo(HaveOccurred())
			Expect(macAddr).To(Equal(netconf.MAC))
			mocked.AssertExpectations(t)
		})
	})

	Context("Checking ReleaseVF function", func() {
		var (
			podifName string
			contID    string
			netconf   *sriovtypes.NetConf
		)

		BeforeEach(func() {
			podifName = "net1"
			contID = "dummycid"
			netconf = &sriovtypes.NetConf{
				Master:      "enp175s0f1",
				DeviceID:    "0000:af:06.0",
				VFID:        0,
				ContIFNames: "net1",
				OrigVfState: sriovtypes.VfState{
					HostIFName: "enp175s6",
				},
			}
		})
		It("Assuming existing interface", func() {
			var targetNetNS ns.NetNS
			targetNetNS, err := testutils.NewNS()
			defer func() {
				if targetNetNS != nil {
					targetNetNS.Close()
				}
			}()
			Expect(err).NotTo(HaveOccurred())
			mocked := &MockNetlinkManager{}
			fakeLink := &FakeLink{netlink.LinkAttrs{Index: 1000, Name: "dummylink"}}

			mocked.On("LinkByName", netconf.ContIFNames).Return(fakeLink, nil)
			mocked.On("LinkSetDown", fakeLink).Return(nil)
			mocked.On("LinkSetName", fakeLink, netconf.OrigVfState.HostIFName).Return(nil)
			mocked.On("LinkSetNsFd", fakeLink, mock.AnythingOfType("int")).Return(nil)
			sm := sriovManager{nLink: mocked}
			err = sm.ReleaseVF(netconf, podifName, contID, targetNetNS)
			Expect(err).NotTo(HaveOccurred())
			mocked.AssertExpectations(t)
		})
	})
	Context("Checking ReleaseVF function - restore config", func() {
		var (
			podifName string
			contID    string
			netconf   *sriovtypes.NetConf
		)

		BeforeEach(func() {
			podifName = "net1"
			contID = "dummycid"
			netconf = &sriovtypes.NetConf{
				Master:      "enp175s0f1",
				DeviceID:    "0000:af:06.0",
				VFID:        0,
				MAC:         "aa:f3:8d:65:1b:d4",
				ContIFNames: "net1",
				OrigVfState: sriovtypes.VfState{
					HostIFName:   "enp175s6",
					EffectiveMAC: "c6:c8:7f:1f:21:90",
				},
			}
		})
		It("Restores Effective MAC address when provided in netconf", func() {
			var targetNetNS ns.NetNS
			targetNetNS, err := testutils.NewNS()
			defer func() {
				if targetNetNS != nil {
					targetNetNS.Close()
				}
			}()
			Expect(err).NotTo(HaveOccurred())
			mocked := &MockNetlinkManager{}
			fakeLink := &FakeLink{netlink.LinkAttrs{Index: 1000, Name: "dummylink"}}

			mocked.On("LinkByName", netconf.ContIFNames).Return(fakeLink, nil)
			mocked.On("LinkSetDown", fakeLink).Return(nil)
			mocked.On("LinkSetName", fakeLink, netconf.OrigVfState.HostIFName).Return(nil)
			mocked.On("LinkSetNsFd", fakeLink, mock.AnythingOfType("int")).Return(nil)
			origEffMac, err := net.ParseMAC(netconf.OrigVfState.EffectiveMAC)
			Expect(err).NotTo(HaveOccurred())
			mocked.On("LinkSetHardwareAddr", fakeLink, origEffMac).Return(nil)
			sm := sriovManager{nLink: mocked}
			err = sm.ReleaseVF(netconf, podifName, contID, targetNetNS)
			Expect(err).NotTo(HaveOccurred())
			mocked.AssertExpectations(t)
		})
	})
	Context("Checking ResetVFConfig function - restore config no user params", func() {
		var (
			netconf *sriovtypes.NetConf
		)

		BeforeEach(func() {
			netconf = &sriovtypes.NetConf{
				Master:      "enp175s0f1",
				DeviceID:    "0000:af:06.0",
				VFID:        0,
				ContIFNames: "net1",
				OrigVfState: sriovtypes.VfState{
					HostIFName: "enp175s6",
				},
			}
		})
		It("Does not change VF config if it wasnt requested to be changed in netconf", func() {
			mocked := &MockNetlinkManager{}
			fakeLink := &FakeLink{netlink.LinkAttrs{Index: 1000, Name: "dummylink"}}

			mocked.On("LinkByName", netconf.Master).Return(fakeLink, nil)
			sm := sriovManager{nLink: mocked}
			err := sm.ResetVFConfig(netconf)
			Expect(err).NotTo(HaveOccurred())
			mocked.AssertExpectations(t)
		})
	})

	Context("Checking ResetVFConfig function - restore config with user params", func() {
		var (
			netconf *sriovtypes.NetConf
		)

		BeforeEach(func() {
			vlan := 6
			vlanQos := 3
			maxTxRate := 4000
			minTxRate := 1000

			netconf = &sriovtypes.NetConf{
				Master:      "enp175s0f1",
				DeviceID:    "0000:af:06.0",
				VFID:        3,
				ContIFNames: "net1",
				MAC:         "d2:fc:22:a7:0d:e8",
				Vlan:        &vlan,
				VlanQoS:     &vlanQos,
				SpoofChk:    "on",
				MaxTxRate:   &maxTxRate,
				MinTxRate:   &minTxRate,
				Trust:       "on",
				LinkState:   "enable",
				OrigVfState: sriovtypes.VfState{
					HostIFName:   "enp175s6",
					SpoofChk:     false,
					AdminMAC:     "aa:f3:8d:65:1b:d4",
					EffectiveMAC: "aa:f3:8d:65:1b:d4",
					Vlan:         1,
					VlanQoS:      1,
					MinTxRate:    0,
					MaxTxRate:    0,
					LinkState:    2, // disable
				},
			}
		})
		It("Restores original VF configurations", func() {
			mocked := &MockNetlinkManager{}
			fakeLink := &FakeLink{netlink.LinkAttrs{Index: 1000, Name: "dummylink"}}

			mocked.On("LinkByName", netconf.Master).Return(fakeLink, nil)
			mocked.On("LinkSetVfVlanQos", fakeLink, netconf.VFID, netconf.OrigVfState.Vlan, netconf.OrigVfState.VlanQoS).Return(nil)
			mocked.On("LinkSetVfSpoofchk", fakeLink, netconf.VFID, netconf.OrigVfState.SpoofChk).Return(nil)
			origMac, err := net.ParseMAC(netconf.OrigVfState.AdminMAC)
			Expect(err).NotTo(HaveOccurred())
			mocked.On("LinkSetVfHardwareAddr", fakeLink, netconf.VFID, origMac).Return(nil)
			mocked.On("LinkSetVfTrust", fakeLink, netconf.VFID, false).Return(nil)
			mocked.On("LinkSetVfRate", fakeLink, netconf.VFID, netconf.OrigVfState.MinTxRate, netconf.OrigVfState.MaxTxRate).Return(nil)
			mocked.On("LinkSetVfState", fakeLink, netconf.VFID, netconf.OrigVfState.LinkState).Return(nil)

			sm := sriovManager{nLink: mocked}
			err = sm.ResetVFConfig(netconf)
			Expect(err).NotTo(HaveOccurred())
			mocked.AssertExpectations(t)
		})
	})
})
