package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
)

const message0210 = "ISO0260000500210B238C4012E8180180000000010000188000000000000001000080311105300012411135208030803539990009000000003165579209013100166=0803101927000000000000001245522850093278402001_____0271234567            12341234484016FIIDTLNETETOPTID019FIIDCALN10100000000010029TERMINALNAMEANDLOCATION00TBID020CLERIDCRTA0808080844012MCHOSTHOST10"

func main() {

	server, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("Eror al iniciar servidor: %+v", err.Error())
		log.Fatal(err)
	}
	fmt.Println("Server on port: 8000")
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Eror al intetnarse comunicar con el servidor: %+v", err.Error())
			log.Fatal(err)
		}
		go func(c net.Conn) {
			for {
				message, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					log.Printf("Error al leer mensaje: %+v", err.Error())
					return
				} else {
					mti := get_Mti(message)
					if mti == "0200" {
						fields := get_DE(message)
						fmt.Printf("\n\nFields recibidos: \n%v", fields)
						es := get_ES(fields)
						var es_07 string
						if es != "-1" {
							es_07 = string(rune(es[len(es)-1]))
						} else {
							es_07 = "0"
						}
						fieldsToServer := get_DE(message0210)
						fieldsToServer[11] = fields[11]
						fieldsToServer[37] = fields[37]
						if es_07 == "1" {
							fmt.Println("\nInicio de llaves") // EA647708D0994C35431605D082451AF7 0102012345678AE00001 E528FD 00 6B7A5FC0
							p63 := "121& 04114! Q100002 02! Q200002 03! C400002 00! EX00068 EA647708D0994C35431605D082451AF70102012345678AE00001E528FD006B7A5FC0"
							fieldsToServer[63] = p63
							fmt.Printf("\nFields To Server: %v", fieldsToServer)
							fmt.Println("\n\nMensaje hacia server: ", get_Message(fieldsToServer, "ISO026000050", "0210", "B238C4012E81801A"))
							conn.Write([]byte(get_Message(fieldsToServer, "ISO026000050", "0210", "B238C4012E81801A")))
						} else {
							fmt.Println("Transaccion normal")
							fmt.Printf("\nFields To Server: %v", fieldsToServer)
							fmt.Println("\n\nMensaje hacia server: ", get_Message(fieldsToServer, "ISO026000050", "0210", "B238C4012E818018"))
							conn.Write([]byte(get_Message(fieldsToServer, "ISO026000050", "0210", "B238C4012E818018")))
						}
						break
					}
					if mti == "0800" {
						fields := get_DE(message)
						fmt.Printf("\n\nFields recibidos: \n%v", fields)
						header := "ISO005000054"
						typeMessage := "0810"
						pbm := "8220000002000000"
						p1 := "0400000000000000"
						p7 := fmt.Sprintf("%02d%02d%02d%02d%02d", int(time.Now().Month()), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
						p11 := fields[11]
						p39 := "00"
						s70 := "301"
						message_0810 := header + typeMessage + pbm + p1 + p7 + p11 + p39 + s70
						fmt.Println("\n\nMensaje hacia server: ", message_0810)
						conn.Write([]byte(message_0810))
					}
				}
			}
			conn.Close()
		}(conn)
	}
}

func get_ES(fields map[int]string) string {
	index := strings.Index(fields[63], "! ES")
	if 5 < len(fields[63]) {
		width := fields[63][index+4 : index+9]
		u64, err := strconv.ParseInt(width, 10, 32)
		if err != nil {
			fmt.Println("Error al parsear")
			fmt.Println(err)
		}
		len := int(u64)
		len += 10
		return fields[63][index : index+len]
	} else {
		return "-1"
	}
}

func get_DE(message string) map[int]string {
	i, err := strconv.ParseUint(message[16:32], 16, 64)
	if err != nil {
		fmt.Printf("%s", err)
	}
	bitmap := fmt.Sprintf("%064b", i)
	init := 32
	fields := make(map[int]string)
	if bitmap[0] == 49 {
		i, err := strconv.ParseUint(message[32:48], 16, 64)
		if err != nil {
			fmt.Printf("%s", err)
		}
		secondBitmap := fmt.Sprintf("%064b", i)
		bitmap = fmt.Sprintf("%s%s", string(bitmap), string(secondBitmap))
	}
	for pos, char := range bitmap {
		if char == 49 {
			switch pos + 1 {
			case 1:
				fields[int(pos+1)] = string(message[init : init+16])
				init += 16
			case 3:
				fields[int(pos+1)] = string(message[init : init+6])
				init += 6
			case 4:
				fields[int(pos+1)] = string(message[init : init+12])
				init += 12
			case 7:
				fields[int(pos+1)] = string(message[init : init+10])
				init += 10
			case 11:
				fields[int(pos+1)] = string(message[init : init+6])
				init += 6
			case 12:
				fields[int(pos+1)] = string(message[init : init+6])
				init += 6
			case 13:
				fields[int(pos+1)] = string(message[init : init+4])
				init += 4
			case 17:
				fields[int(pos+1)] = string(message[init : init+4])
				init += 4
			case 18:
				fields[int(pos+1)] = string(message[init : init+4])
				init += 4
			case 22:
				fields[int(pos+1)] = string(message[init : init+3])
				init += 3
			case 25:
				fields[int(pos+1)] = string(message[init : init+2])
				init += 2
			case 32:
				fields[int(pos+1)] = string(message[init : init+11])
				init += 11
			case 35:
				width := message[init : init+2]
				u64, err := strconv.ParseInt(width, 10, 32)
				if err != nil {
					fmt.Println(err)
				}
				len := int(u64)
				len += 2 + 16
				fields[int(pos+1)] = string(message[init : init+len])
				init += len
			case 37:
				fields[int(pos+1)] = string(message[init : init+12])
				init += 12
			case 38:
				fields[int(pos+1)] = string(message[init : init+6])
				init += 6
			case 39:
				fields[int(pos+1)] = string(message[init : init+2])
				init += 2
			case 41:
				fields[int(pos+1)] = string(message[init : init+16])
				init += 16
			case 42:
				fields[int(pos+1)] = string(message[init : init+15])
				init += 15
			case 43:
				fields[int(pos+1)] = string(message[init : init+40])
				init += 40
			case 44:
				fields[int(pos+1)] = string(message[init : init+4])
				init += 4
			case 48, 61, 63, 120, 121, 125, 126:
				width := message[init : init+3]
				u64, err := strconv.ParseInt(width, 10, 32)
				if err != nil {
					fmt.Println("Error al parsear")
					fmt.Println(err)
				}
				len := int(u64)
				len += 3
				fields[int(pos+1)] = string(message[init : init+len])
				init += len
			case 49:
				fields[int(pos+1)] = string(message[init : init+3])
				init += 3
			case 60:
				fields[int(pos+1)] = string(message[init : init+19])
				init += 19
			case 100:
				width := message[init : init+2]
				u64, err := strconv.ParseInt(width, 10, 32)
				if err != nil {
					fmt.Println(err)
				}
				len := int(u64)
				len += 2
				fields[int(pos+1)] = string(message[init : init+len])
				init += len
			}
		}
	}
	return fields
}

func get_Message(dataElements map[int]string, header string, mti string, bitmap string) string {
	var fields []string
	keys := make([]int, 0, len(dataElements))
	for k := range dataElements {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fields = append(fields, dataElements[key])
	}
	de := strings.Join(fields, "")
	message := header + mti + bitmap + de
	return message
}
func get_Mti(message string) string {
	mti := message[12:16]
	return mti
}
