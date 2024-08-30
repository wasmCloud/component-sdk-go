package net

import (
	"net"
	"net/netip"
	"time"

	"tinygo.org/x/drivers/netdev"
	"tinygo.org/x/drivers/netlink"
)

type Config struct{}

var (
	_ netlink.Netlinker = &Device{}
	_ netdev.Netdever   = &Device{}
)

type Device struct {
	cfg *Config
}

func NewDevice(cfg *Config) *Device {
	return &Device{
		cfg: cfg,
	}
}

func (d *Device) NetConnect(params *netlink.ConnectParams) error {
	panic("not implemented")
}

func (d *Device) NetDisconnect() {
	panic("not implemented")
}

func (d *Device) NetNotify(cb func(netlink.Event)) {
	panic("not implemented")
}

func (d *Device) GetHostByName(name string) (netip.Addr, error) {
	panic("not implemented")
}

func (d *Device) GetHardwareAddr() (net.HardwareAddr, error) {
	panic("not implemented")
}

func (d *Device) Addr() (netip.Addr, error) {
	panic("not implemented")
}

func (d *Device) Socket(domain int, stype int, protocol int) (int, error) {
	panic("not implemented")
}

func (d *Device) Bind(sockfd int, ip netip.AddrPort) error {
	panic("not implemented")
}

func (d *Device) Connect(sockfd int, host string, ip netip.AddrPort) error {
	panic("not implemented")
}

func (d *Device) Listen(sockfd int, backlog int) error {
	panic("not implemented")
}

func (d *Device) Accept(sockfd int) (int, netip.AddrPort, error) {
	panic("not implemented")
}

func (d *Device) sendChunk(sockfd int, buf []byte, deadline time.Time) (int, error) {
	panic("not implemented")
}

func (d *Device) Send(sockfd int, buf []byte, flags int, deadline time.Time) (int, error) {
	panic("not implemented")
}

func (d *Device) Recv(sockfd int, buf []byte, flags int, deadline time.Time) (int, error) {
	panic("not implemented")
}

func (d *Device) Close(sockfd int) error {
	panic("not implemented")
}

func (d *Device) SetSockOpt(sockfd int, level int, opt int, value interface{}) error {
	panic("not implemented")
}

// Connected checks if there is communication with the ESP8266/ESP32.
func (d *Device) Connected() bool {
	panic("not implemented")
}

// Write raw bytes to the UART.
func (d *Device) Write(b []byte) (n int, err error) {
	panic("not implemented")
}

// Read raw bytes from the UART.
func (d *Device) Read(b []byte) (n int, err error) {
	panic("not implemented")
}

// Execute sends an AT command to the ESP8266/ESP32.
func (d Device) Execute(cmd string) error {
	panic("not implemented")
}

// Query sends an AT command to the ESP8266/ESP32 that returns the
// current value for some configuration parameter.
func (d Device) Query(cmd string) (string, error) {
	panic("not implemented")
}

// Set sends an AT command with params to the ESP8266/ESP32 for a
// configuration value to be set.
func (d Device) Set(cmd, params string) error {
	panic("not implemented")
}

// Version returns the ESP8266/ESP32 firmware version info.
func (d Device) Version() []byte {
	panic("not implemented")
}

// Echo sets the ESP8266/ESP32 echo setting.
func (d Device) Echo(set bool) {
	panic("not implemented")
}

// Reset restarts the ESP8266/ESP32 firmware. Due to how the baud rate changes,
// this messes up communication with the ESP8266/ESP32 module. So make sure you know
// what you are doing when you call this.
func (d Device) Reset() {
	panic("not implemented")
}

// ReadSocket returns the data that has already been read in from the responses.
func (d *Device) ReadSocket(b []byte) (n int, err error) {
	panic("not implemented")
}

// Response gets the next response bytes from the ESP8266/ESP32.
// The call will retry for up to timeout milliseconds before returning nothing.
func (d *Device) Response(timeout int) ([]byte, error) {
	panic("not implemented")
}

func (d *Device) parseIPD(end int) error {
	panic("not implemented")
}

// IsSocketDataAvailable returns of there is socket data available
func (d *Device) IsSocketDataAvailable() bool {
	panic("not implemented")
}

func EnableSockets() {
	cfg := Config{}
	dev := NewDevice(&cfg)
	netdev.UseNetdev(dev)
}
