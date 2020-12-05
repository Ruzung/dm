/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_864 struct{}

var Dm_build_865 = &dm_build_864{}

func (Dm_build_867 *dm_build_864) Dm_build_866(dm_build_868 []byte, dm_build_869 int, dm_build_870 byte) int {
	dm_build_868[dm_build_869] = dm_build_870
	return 1
}

func (Dm_build_872 *dm_build_864) Dm_build_871(dm_build_873 []byte, dm_build_874 int, dm_build_875 int8) int {
	dm_build_873[dm_build_874] = byte(dm_build_875)
	return 1
}

func (Dm_build_877 *dm_build_864) Dm_build_876(dm_build_878 []byte, dm_build_879 int, dm_build_880 int16) int {
	dm_build_878[dm_build_879] = byte(dm_build_880)
	dm_build_879++
	dm_build_878[dm_build_879] = byte(dm_build_880 >> 8)
	return 2
}

func (Dm_build_882 *dm_build_864) Dm_build_881(dm_build_883 []byte, dm_build_884 int, dm_build_885 int32) int {
	dm_build_883[dm_build_884] = byte(dm_build_885)
	dm_build_884++
	dm_build_883[dm_build_884] = byte(dm_build_885 >> 8)
	dm_build_884++
	dm_build_883[dm_build_884] = byte(dm_build_885 >> 16)
	dm_build_884++
	dm_build_883[dm_build_884] = byte(dm_build_885 >> 24)
	dm_build_884++
	return 4
}

func (Dm_build_887 *dm_build_864) Dm_build_886(dm_build_888 []byte, dm_build_889 int, dm_build_890 int64) int {
	dm_build_888[dm_build_889] = byte(dm_build_890)
	dm_build_889++
	dm_build_888[dm_build_889] = byte(dm_build_890 >> 8)
	dm_build_889++
	dm_build_888[dm_build_889] = byte(dm_build_890 >> 16)
	dm_build_889++
	dm_build_888[dm_build_889] = byte(dm_build_890 >> 24)
	dm_build_889++
	dm_build_888[dm_build_889] = byte(dm_build_890 >> 32)
	dm_build_889++
	dm_build_888[dm_build_889] = byte(dm_build_890 >> 40)
	dm_build_889++
	dm_build_888[dm_build_889] = byte(dm_build_890 >> 48)
	dm_build_889++
	dm_build_888[dm_build_889] = byte(dm_build_890 >> 56)
	return 8
}

func (Dm_build_892 *dm_build_864) Dm_build_891(dm_build_893 []byte, dm_build_894 int, dm_build_895 float32) int {
	return Dm_build_892.Dm_build_911(dm_build_893, dm_build_894, math.Float32bits(dm_build_895))
}

func (Dm_build_897 *dm_build_864) Dm_build_896(dm_build_898 []byte, dm_build_899 int, dm_build_900 float64) int {
	return Dm_build_897.Dm_build_916(dm_build_898, dm_build_899, math.Float64bits(dm_build_900))
}

func (Dm_build_902 *dm_build_864) Dm_build_901(dm_build_903 []byte, dm_build_904 int, dm_build_905 uint8) int {
	dm_build_903[dm_build_904] = byte(dm_build_905)
	return 1
}

func (Dm_build_907 *dm_build_864) Dm_build_906(dm_build_908 []byte, dm_build_909 int, dm_build_910 uint16) int {
	dm_build_908[dm_build_909] = byte(dm_build_910)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 8)
	return 2
}

func (Dm_build_912 *dm_build_864) Dm_build_911(dm_build_913 []byte, dm_build_914 int, dm_build_915 uint32) int {
	dm_build_913[dm_build_914] = byte(dm_build_915)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 8)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 16)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 24)
	return 3
}

func (Dm_build_917 *dm_build_864) Dm_build_916(dm_build_918 []byte, dm_build_919 int, dm_build_920 uint64) int {
	dm_build_918[dm_build_919] = byte(dm_build_920)
	dm_build_919++
	dm_build_918[dm_build_919] = byte(dm_build_920 >> 8)
	dm_build_919++
	dm_build_918[dm_build_919] = byte(dm_build_920 >> 16)
	dm_build_919++
	dm_build_918[dm_build_919] = byte(dm_build_920 >> 24)
	dm_build_919++
	dm_build_918[dm_build_919] = byte(dm_build_920 >> 32)
	dm_build_919++
	dm_build_918[dm_build_919] = byte(dm_build_920 >> 40)
	dm_build_919++
	dm_build_918[dm_build_919] = byte(dm_build_920 >> 48)
	dm_build_919++
	dm_build_918[dm_build_919] = byte(dm_build_920 >> 56)
	return 3
}

func (Dm_build_922 *dm_build_864) Dm_build_921(dm_build_923 []byte, dm_build_924 int, dm_build_925 []byte, dm_build_926 int, dm_build_927 int) int {
	copy(dm_build_923[dm_build_924:dm_build_924+dm_build_927], dm_build_925[dm_build_926:dm_build_926+dm_build_927])
	return dm_build_927
}

func (Dm_build_929 *dm_build_864) Dm_build_928(dm_build_930 []byte, dm_build_931 int, dm_build_932 []byte, dm_build_933 int, dm_build_934 int) int {
	dm_build_931 += Dm_build_929.Dm_build_911(dm_build_930, dm_build_931, uint32(dm_build_934))
	return 4 + Dm_build_929.Dm_build_921(dm_build_930, dm_build_931, dm_build_932, dm_build_933, dm_build_934)
}

func (Dm_build_936 *dm_build_864) Dm_build_935(dm_build_937 []byte, dm_build_938 int, dm_build_939 []byte, dm_build_940 int, dm_build_941 int) int {
	dm_build_938 += Dm_build_936.Dm_build_906(dm_build_937, dm_build_938, uint16(dm_build_941))
	return 2 + Dm_build_936.Dm_build_921(dm_build_937, dm_build_938, dm_build_939, dm_build_940, dm_build_941)
}

func (Dm_build_943 *dm_build_864) Dm_build_942(dm_build_944 []byte, dm_build_945 int, dm_build_946 string, dm_build_947 string, dm_build_948 *DmConnection) int {
	dm_build_949 := Dm_build_943.Dm_build_1075(dm_build_946, dm_build_947, dm_build_948)
	dm_build_945 += Dm_build_943.Dm_build_911(dm_build_944, dm_build_945, uint32(len(dm_build_949)))
	return 4 + Dm_build_943.Dm_build_921(dm_build_944, dm_build_945, dm_build_949, 0, len(dm_build_949))
}

func (Dm_build_951 *dm_build_864) Dm_build_950(dm_build_952 []byte, dm_build_953 int, dm_build_954 string, dm_build_955 string, dm_build_956 *DmConnection) int {
	dm_build_957 := Dm_build_951.Dm_build_1075(dm_build_954, dm_build_955, dm_build_956)

	dm_build_953 += Dm_build_951.Dm_build_906(dm_build_952, dm_build_953, uint16(len(dm_build_957)))
	return 2 + Dm_build_951.Dm_build_921(dm_build_952, dm_build_953, dm_build_957, 0, len(dm_build_957))
}

func (Dm_build_959 *dm_build_864) Dm_build_958(dm_build_960 []byte, dm_build_961 int) byte {
	return dm_build_960[dm_build_961]
}

func (Dm_build_963 *dm_build_864) Dm_build_962(dm_build_964 []byte, dm_build_965 int) int16 {
	var dm_build_966 int16
	dm_build_966 = int16(dm_build_964[dm_build_965] & 0xff)
	dm_build_965++
	dm_build_966 |= int16(dm_build_964[dm_build_965]&0xff) << 8
	return dm_build_966
}

func (Dm_build_968 *dm_build_864) Dm_build_967(dm_build_969 []byte, dm_build_970 int) int32 {
	var dm_build_971 int32
	dm_build_971 = int32(dm_build_969[dm_build_970] & 0xff)
	dm_build_970++
	dm_build_971 |= int32(dm_build_969[dm_build_970]&0xff) << 8
	dm_build_970++
	dm_build_971 |= int32(dm_build_969[dm_build_970]&0xff) << 16
	dm_build_970++
	dm_build_971 |= int32(dm_build_969[dm_build_970]&0xff) << 24
	return dm_build_971
}

func (Dm_build_973 *dm_build_864) Dm_build_972(dm_build_974 []byte, dm_build_975 int) int64 {
	var dm_build_976 int64
	dm_build_976 = int64(dm_build_974[dm_build_975] & 0xff)
	dm_build_975++
	dm_build_976 |= int64(dm_build_974[dm_build_975]&0xff) << 8
	dm_build_975++
	dm_build_976 |= int64(dm_build_974[dm_build_975]&0xff) << 16
	dm_build_975++
	dm_build_976 |= int64(dm_build_974[dm_build_975]&0xff) << 24
	dm_build_975++
	dm_build_976 |= int64(dm_build_974[dm_build_975]&0xff) << 32
	dm_build_975++
	dm_build_976 |= int64(dm_build_974[dm_build_975]&0xff) << 40
	dm_build_975++
	dm_build_976 |= int64(dm_build_974[dm_build_975]&0xff) << 48
	dm_build_975++
	dm_build_976 |= int64(dm_build_974[dm_build_975]&0xff) << 56
	return dm_build_976
}

func (Dm_build_978 *dm_build_864) Dm_build_977(dm_build_979 []byte, dm_build_980 int) float32 {
	return math.Float32frombits(Dm_build_978.Dm_build_994(dm_build_979, dm_build_980))
}

func (Dm_build_982 *dm_build_864) Dm_build_981(dm_build_983 []byte, dm_build_984 int) float64 {
	return math.Float64frombits(Dm_build_982.Dm_build_999(dm_build_983, dm_build_984))
}

func (Dm_build_986 *dm_build_864) Dm_build_985(dm_build_987 []byte, dm_build_988 int) uint8 {
	return uint8(dm_build_987[dm_build_988] & 0xff)
}

func (Dm_build_990 *dm_build_864) Dm_build_989(dm_build_991 []byte, dm_build_992 int) uint16 {
	var dm_build_993 uint16
	dm_build_993 = uint16(dm_build_991[dm_build_992] & 0xff)
	dm_build_992++
	dm_build_993 |= uint16(dm_build_991[dm_build_992]&0xff) << 8
	return dm_build_993
}

func (Dm_build_995 *dm_build_864) Dm_build_994(dm_build_996 []byte, dm_build_997 int) uint32 {
	var dm_build_998 uint32
	dm_build_998 = uint32(dm_build_996[dm_build_997] & 0xff)
	dm_build_997++
	dm_build_998 |= uint32(dm_build_996[dm_build_997]&0xff) << 8
	dm_build_997++
	dm_build_998 |= uint32(dm_build_996[dm_build_997]&0xff) << 16
	dm_build_997++
	dm_build_998 |= uint32(dm_build_996[dm_build_997]&0xff) << 24
	return dm_build_998
}

func (Dm_build_1000 *dm_build_864) Dm_build_999(dm_build_1001 []byte, dm_build_1002 int) uint64 {
	var dm_build_1003 uint64
	dm_build_1003 = uint64(dm_build_1001[dm_build_1002] & 0xff)
	dm_build_1002++
	dm_build_1003 |= uint64(dm_build_1001[dm_build_1002]&0xff) << 8
	dm_build_1002++
	dm_build_1003 |= uint64(dm_build_1001[dm_build_1002]&0xff) << 16
	dm_build_1002++
	dm_build_1003 |= uint64(dm_build_1001[dm_build_1002]&0xff) << 24
	dm_build_1002++
	dm_build_1003 |= uint64(dm_build_1001[dm_build_1002]&0xff) << 32
	dm_build_1002++
	dm_build_1003 |= uint64(dm_build_1001[dm_build_1002]&0xff) << 40
	dm_build_1002++
	dm_build_1003 |= uint64(dm_build_1001[dm_build_1002]&0xff) << 48
	dm_build_1002++
	dm_build_1003 |= uint64(dm_build_1001[dm_build_1002]&0xff) << 56
	return dm_build_1003
}

func (Dm_build_1005 *dm_build_864) Dm_build_1004(dm_build_1006 []byte, dm_build_1007 int) []byte {
	dm_build_1008 := Dm_build_1005.Dm_build_994(dm_build_1006, dm_build_1007)
	return dm_build_1006[dm_build_1007+4 : dm_build_1007+4+int(dm_build_1008)]

}

func (Dm_build_1010 *dm_build_864) Dm_build_1009(dm_build_1011 []byte, dm_build_1012 int) []byte {
	dm_build_1013 := Dm_build_1010.Dm_build_989(dm_build_1011, dm_build_1012)
	return dm_build_1011[dm_build_1012+2 : dm_build_1012+2+int(dm_build_1013)]

}

func (Dm_build_1015 *dm_build_864) Dm_build_1014(dm_build_1016 []byte, dm_build_1017 int, dm_build_1018 int) []byte {
	return dm_build_1016[dm_build_1017 : dm_build_1017+dm_build_1018]

}

func (Dm_build_1020 *dm_build_864) Dm_build_1019(dm_build_1021 []byte, dm_build_1022 int, dm_build_1023 int, dm_build_1024 string, dm_build_1025 *DmConnection) string {
	return Dm_build_1020.Dm_build_1112(dm_build_1021[dm_build_1022:dm_build_1022+dm_build_1023], dm_build_1024, dm_build_1025)
}

func (Dm_build_1027 *dm_build_864) Dm_build_1026(dm_build_1028 []byte, dm_build_1029 int, dm_build_1030 string, dm_build_1031 *DmConnection) string {
	dm_build_1032 := Dm_build_1027.Dm_build_994(dm_build_1028, dm_build_1029)
	dm_build_1029 += 4
	return Dm_build_1027.Dm_build_1019(dm_build_1028, dm_build_1029, int(dm_build_1032), dm_build_1030, dm_build_1031)
}

func (Dm_build_1034 *dm_build_864) Dm_build_1033(dm_build_1035 []byte, dm_build_1036 int, dm_build_1037 string, dm_build_1038 *DmConnection) string {
	dm_build_1039 := Dm_build_1034.Dm_build_989(dm_build_1035, dm_build_1036)
	dm_build_1036 += 2
	return Dm_build_1034.Dm_build_1019(dm_build_1035, dm_build_1036, int(dm_build_1039), dm_build_1037, dm_build_1038)
}

func (Dm_build_1041 *dm_build_864) Dm_build_1040(dm_build_1042 byte) []byte {
	return []byte{dm_build_1042}
}

func (Dm_build_1044 *dm_build_864) Dm_build_1043(dm_build_1045 int16) []byte {
	return []byte{byte(dm_build_1045), byte(dm_build_1045 >> 8)}
}

func (Dm_build_1047 *dm_build_864) Dm_build_1046(dm_build_1048 int32) []byte {
	return []byte{byte(dm_build_1048), byte(dm_build_1048 >> 8), byte(dm_build_1048 >> 16), byte(dm_build_1048 >> 24)}
}

func (Dm_build_1050 *dm_build_864) Dm_build_1049(dm_build_1051 int64) []byte {
	return []byte{byte(dm_build_1051), byte(dm_build_1051 >> 8), byte(dm_build_1051 >> 16), byte(dm_build_1051 >> 24), byte(dm_build_1051 >> 32),
		byte(dm_build_1051 >> 40), byte(dm_build_1051 >> 48), byte(dm_build_1051 >> 56)}
}

func (Dm_build_1053 *dm_build_864) Dm_build_1052(dm_build_1054 float32) []byte {
	return Dm_build_1053.Dm_build_1064(math.Float32bits(dm_build_1054))
}

func (Dm_build_1056 *dm_build_864) Dm_build_1055(dm_build_1057 float64) []byte {
	return Dm_build_1056.Dm_build_1067(math.Float64bits(dm_build_1057))
}

func (Dm_build_1059 *dm_build_864) Dm_build_1058(dm_build_1060 uint8) []byte {
	return []byte{byte(dm_build_1060)}
}

func (Dm_build_1062 *dm_build_864) Dm_build_1061(dm_build_1063 uint16) []byte {
	return []byte{byte(dm_build_1063), byte(dm_build_1063 >> 8)}
}

func (Dm_build_1065 *dm_build_864) Dm_build_1064(dm_build_1066 uint32) []byte {
	return []byte{byte(dm_build_1066), byte(dm_build_1066 >> 8), byte(dm_build_1066 >> 16), byte(dm_build_1066 >> 24)}
}

func (Dm_build_1068 *dm_build_864) Dm_build_1067(dm_build_1069 uint64) []byte {
	return []byte{byte(dm_build_1069), byte(dm_build_1069 >> 8), byte(dm_build_1069 >> 16), byte(dm_build_1069 >> 24), byte(dm_build_1069 >> 32), byte(dm_build_1069 >> 40), byte(dm_build_1069 >> 48), byte(dm_build_1069 >> 56)}
}

func (Dm_build_1071 *dm_build_864) Dm_build_1070(dm_build_1072 []byte, dm_build_1073 string, dm_build_1074 *DmConnection) []byte {
	if dm_build_1073 == "UTF-8" {
		return dm_build_1072
	}

	if dm_build_1074 == nil {
		if e := dm_build_1117(dm_build_1073); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1072), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1074.encodeBuffer == nil {
		dm_build_1074.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1074.encode = dm_build_1117(dm_build_1074.getServerEncoding())
		dm_build_1074.transformReaderDst = make([]byte, 4096)
		dm_build_1074.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1074.encode; e != nil {

		dm_build_1074.encodeBuffer.Reset()

		n, err := dm_build_1074.encodeBuffer.ReadFrom(
			Dm_build_1131(bytes.NewReader(dm_build_1072), e.NewEncoder(), dm_build_1074.transformReaderDst, dm_build_1074.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1074.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1076 *dm_build_864) Dm_build_1075(dm_build_1077 string, dm_build_1078 string, dm_build_1079 *DmConnection) []byte {
	return Dm_build_1076.Dm_build_1070([]byte(dm_build_1077), dm_build_1078, dm_build_1079)
}

func (Dm_build_1081 *dm_build_864) Dm_build_1080(dm_build_1082 []byte) byte {
	return Dm_build_1081.Dm_build_958(dm_build_1082, 0)
}

func (Dm_build_1084 *dm_build_864) Dm_build_1083(dm_build_1085 []byte) int16 {
	return Dm_build_1084.Dm_build_962(dm_build_1085, 0)
}

func (Dm_build_1087 *dm_build_864) Dm_build_1086(dm_build_1088 []byte) int32 {
	return Dm_build_1087.Dm_build_967(dm_build_1088, 0)
}

func (Dm_build_1090 *dm_build_864) Dm_build_1089(dm_build_1091 []byte) int64 {
	return Dm_build_1090.Dm_build_972(dm_build_1091, 0)
}

func (Dm_build_1093 *dm_build_864) Dm_build_1092(dm_build_1094 []byte) float32 {
	return Dm_build_1093.Dm_build_977(dm_build_1094, 0)
}

func (Dm_build_1096 *dm_build_864) Dm_build_1095(dm_build_1097 []byte) float64 {
	return Dm_build_1096.Dm_build_981(dm_build_1097, 0)
}

func (Dm_build_1099 *dm_build_864) Dm_build_1098(dm_build_1100 []byte) uint8 {
	return Dm_build_1099.Dm_build_985(dm_build_1100, 0)
}

func (Dm_build_1102 *dm_build_864) Dm_build_1101(dm_build_1103 []byte) uint16 {
	return Dm_build_1102.Dm_build_989(dm_build_1103, 0)
}

func (Dm_build_1105 *dm_build_864) Dm_build_1104(dm_build_1106 []byte) uint32 {
	return Dm_build_1105.Dm_build_994(dm_build_1106, 0)
}

func (Dm_build_1108 *dm_build_864) Dm_build_1107(dm_build_1109 []byte, dm_build_1110 string, dm_build_1111 *DmConnection) []byte {
	if dm_build_1110 == "UTF-8" {
		return dm_build_1109
	}

	if dm_build_1111 == nil {
		if e := dm_build_1117(dm_build_1110); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1109), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1111.encodeBuffer == nil {
		dm_build_1111.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1111.encode = dm_build_1117(dm_build_1111.getServerEncoding())
		dm_build_1111.transformReaderDst = make([]byte, 4096)
		dm_build_1111.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1111.encode; e != nil {

		dm_build_1111.encodeBuffer.Reset()

		n, err := dm_build_1111.encodeBuffer.ReadFrom(
			Dm_build_1131(bytes.NewReader(dm_build_1109), e.NewDecoder(), dm_build_1111.transformReaderDst, dm_build_1111.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1111.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1113 *dm_build_864) Dm_build_1112(dm_build_1114 []byte, dm_build_1115 string, dm_build_1116 *DmConnection) string {
	return string(Dm_build_1113.Dm_build_1107(dm_build_1114, dm_build_1115, dm_build_1116))
}

func dm_build_1117(dm_build_1118 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1118); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1119 struct {
	dm_build_1120 io.Reader
	dm_build_1121 transform.Transformer
	dm_build_1122 error

	dm_build_1123                []byte
	dm_build_1124, dm_build_1125 int

	dm_build_1126                []byte
	dm_build_1127, dm_build_1128 int

	dm_build_1129 bool
}

const dm_build_1130 = 4096

func Dm_build_1131(dm_build_1132 io.Reader, dm_build_1133 transform.Transformer, dm_build_1134 []byte, dm_build_1135 []byte) *Dm_build_1119 {
	dm_build_1133.Reset()
	return &Dm_build_1119{
		dm_build_1120: dm_build_1132,
		dm_build_1121: dm_build_1133,
		dm_build_1123: dm_build_1134,
		dm_build_1126: dm_build_1135,
	}
}

func (dm_build_1137 *Dm_build_1119) Read(dm_build_1138 []byte) (int, error) {
	dm_build_1139, dm_build_1140 := 0, error(nil)
	for {

		if dm_build_1137.dm_build_1124 != dm_build_1137.dm_build_1125 {
			dm_build_1139 = copy(dm_build_1138, dm_build_1137.dm_build_1123[dm_build_1137.dm_build_1124:dm_build_1137.dm_build_1125])
			dm_build_1137.dm_build_1124 += dm_build_1139
			if dm_build_1137.dm_build_1124 == dm_build_1137.dm_build_1125 && dm_build_1137.dm_build_1129 {
				return dm_build_1139, dm_build_1137.dm_build_1122
			}
			return dm_build_1139, nil
		} else if dm_build_1137.dm_build_1129 {
			return 0, dm_build_1137.dm_build_1122
		}

		if dm_build_1137.dm_build_1127 != dm_build_1137.dm_build_1128 || dm_build_1137.dm_build_1122 != nil {
			dm_build_1137.dm_build_1124 = 0
			dm_build_1137.dm_build_1125, dm_build_1139, dm_build_1140 = dm_build_1137.dm_build_1121.Transform(dm_build_1137.dm_build_1123, dm_build_1137.dm_build_1126[dm_build_1137.dm_build_1127:dm_build_1137.dm_build_1128], dm_build_1137.dm_build_1122 == io.EOF)
			dm_build_1137.dm_build_1127 += dm_build_1139

			switch {
			case dm_build_1140 == nil:
				if dm_build_1137.dm_build_1127 != dm_build_1137.dm_build_1128 {
					dm_build_1137.dm_build_1122 = nil
				}

				dm_build_1137.dm_build_1129 = dm_build_1137.dm_build_1122 != nil
				continue
			case dm_build_1140 == transform.ErrShortDst && (dm_build_1137.dm_build_1125 != 0 || dm_build_1139 != 0):

				continue
			case dm_build_1140 == transform.ErrShortSrc && dm_build_1137.dm_build_1128-dm_build_1137.dm_build_1127 != len(dm_build_1137.dm_build_1126) && dm_build_1137.dm_build_1122 == nil:

			default:
				dm_build_1137.dm_build_1129 = true

				if dm_build_1137.dm_build_1122 == nil || dm_build_1137.dm_build_1122 == io.EOF {
					dm_build_1137.dm_build_1122 = dm_build_1140
				}
				continue
			}
		}

		if dm_build_1137.dm_build_1127 != 0 {
			dm_build_1137.dm_build_1127, dm_build_1137.dm_build_1128 = 0, copy(dm_build_1137.dm_build_1126, dm_build_1137.dm_build_1126[dm_build_1137.dm_build_1127:dm_build_1137.dm_build_1128])
		}
		dm_build_1139, dm_build_1137.dm_build_1122 = dm_build_1137.dm_build_1120.Read(dm_build_1137.dm_build_1126[dm_build_1137.dm_build_1128:])
		dm_build_1137.dm_build_1128 += dm_build_1139
	}
}
