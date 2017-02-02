package gameboygo
type flag uint8
const(
	ZERO 		flag = 0x80
	SUBTRACT 	flag = 0x40
	HALFCARRY	flag = 0x20
	CARRY		flag = 0X10
)
var ram [0x10000]uint8
var Rom1 		= ram[0x0000:0x4000]
var RomN 		= ram[0x4000:0x8000]
var Vram 		= ram[0x8000:0xA000]
var CartRam 	= ram[0xA000:0xC000]
var Wram 		= ram[0xC000:0xE000]
var EchoRam 	= ram[0xC000:0xDE00]
var Oram 		= ram[0xFE00:0xFEA0]
var Io 			= ram[0xFF00:0xFF80]
var Hram 		= ram[0xFF80:0xFFFF]
var InterruptReg = &ram[0xFFFF]

/*
Writing to this register launches a DMA transfer from ROM or RAM to OAM memory (sprite attribute table). The written value specifies the transfer source address divided by 100h, ie. source & destination are:
  Source:      XX00-XX9F   ;XX in range from 00-F1h
  Destination: FE00-FE9F
*/
var Dma = &ram[0xFF46] 
func setFlags(f flag){
	*InterruptReg |= uint8(f)
}
func clearFlags(f flag) {
	*InterruptReg = *InterruptReg &^ uint8(f)
}
func writeByte(addr uint16, b uint8) bool{
	if addr < 0x8000 {
		return false
	}
	ram[addr] = b
	return true
}
func readByte(addr uint16) uint8{
	return ram[addr]
}
func read16bits(addr uint16) uint16{
	return uint16(uint16(ram[addr]) | (uint16(ram[addr+1]) << 8))
}
func write16bits(addr uint16, data uint16) {
	writeByte(addr, uint8(data & 0x00FF))
	writeByte(addr+1,uint8((data & 0xFF00) >> 8))
}