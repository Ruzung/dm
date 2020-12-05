/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1141 struct {
	dm_build_1142 *list.List
	dm_build_1143 *dm_build_1195
	dm_build_1144 int
}

func Dm_build_1145() *Dm_build_1141 {
	return &Dm_build_1141{
		dm_build_1142: list.New(),
		dm_build_1144: 0,
	}
}

func (dm_build_1147 *Dm_build_1141) Dm_build_1146() int {
	return dm_build_1147.dm_build_1144
}

func (dm_build_1149 *Dm_build_1141) Dm_build_1148(dm_build_1150 *Dm_build_1219, dm_build_1151 int) int {
	var dm_build_1152 = 0
	var dm_build_1153 = 0
	for dm_build_1152 < dm_build_1151 && dm_build_1149.dm_build_1143 != nil {
		dm_build_1153 = dm_build_1149.dm_build_1143.dm_build_1203(dm_build_1150, dm_build_1151-dm_build_1152)
		if dm_build_1149.dm_build_1143.dm_build_1198 == 0 {
			dm_build_1149.dm_build_1185()
		}
		dm_build_1152 += dm_build_1153
		dm_build_1149.dm_build_1144 -= dm_build_1153
	}
	return dm_build_1152
}

func (dm_build_1155 *Dm_build_1141) Dm_build_1154(dm_build_1156 []byte, dm_build_1157 int, dm_build_1158 int) int {
	var dm_build_1159 = 0
	var dm_build_1160 = 0
	for dm_build_1159 < dm_build_1158 && dm_build_1155.dm_build_1143 != nil {
		dm_build_1160 = dm_build_1155.dm_build_1143.dm_build_1207(dm_build_1156, dm_build_1157, dm_build_1158-dm_build_1159)
		if dm_build_1155.dm_build_1143.dm_build_1198 == 0 {
			dm_build_1155.dm_build_1185()
		}
		dm_build_1159 += dm_build_1160
		dm_build_1155.dm_build_1144 -= dm_build_1160
		dm_build_1157 += dm_build_1160
	}
	return dm_build_1159
}

func (dm_build_1162 *Dm_build_1141) Dm_build_1161(dm_build_1163 io.Writer, dm_build_1164 int) int {
	var dm_build_1165 = 0
	var dm_build_1166 = 0
	for dm_build_1165 < dm_build_1164 && dm_build_1162.dm_build_1143 != nil {
		dm_build_1166 = dm_build_1162.dm_build_1143.dm_build_1212(dm_build_1163, dm_build_1164-dm_build_1165)
		if dm_build_1162.dm_build_1143.dm_build_1198 == 0 {
			dm_build_1162.dm_build_1185()
		}
		dm_build_1165 += dm_build_1166
		dm_build_1162.dm_build_1144 -= dm_build_1166
	}
	return dm_build_1165
}

func (dm_build_1168 *Dm_build_1141) Dm_build_1167(dm_build_1169 []byte, dm_build_1170 int, dm_build_1171 int) {
	if dm_build_1171 == 0 {
		return
	}
	var dm_build_1172 = dm_build_1199(dm_build_1169, dm_build_1170, dm_build_1171)
	if dm_build_1168.dm_build_1143 == nil {
		dm_build_1168.dm_build_1143 = dm_build_1172
	} else {
		dm_build_1168.dm_build_1142.PushBack(dm_build_1172)
	}
	dm_build_1168.dm_build_1144 += dm_build_1171
}

func (dm_build_1174 *Dm_build_1141) dm_build_1173(dm_build_1175 int) byte {
	var dm_build_1176 = dm_build_1175
	var dm_build_1177 = dm_build_1174.dm_build_1143
	for dm_build_1176 > 0 && dm_build_1177 != nil {
		if dm_build_1177.dm_build_1198 == 0 {
			continue
		}
		if dm_build_1176 > dm_build_1177.dm_build_1198-1 {
			dm_build_1176 -= dm_build_1177.dm_build_1198
			dm_build_1177 = dm_build_1174.dm_build_1142.Front().Value.(*dm_build_1195)
		} else {
			break
		}
	}
	return dm_build_1177.dm_build_1216(dm_build_1176)
}
func (dm_build_1179 *Dm_build_1141) Dm_build_1178(dm_build_1180 *Dm_build_1141) {
	if dm_build_1180.dm_build_1144 == 0 {
		return
	}
	var dm_build_1181 = dm_build_1180.dm_build_1143
	for dm_build_1181 != nil {
		dm_build_1179.dm_build_1182(dm_build_1181)
		dm_build_1180.dm_build_1185()
		dm_build_1181 = dm_build_1180.dm_build_1143
	}
	dm_build_1180.dm_build_1144 = 0
}
func (dm_build_1183 *Dm_build_1141) dm_build_1182(dm_build_1184 *dm_build_1195) {
	if dm_build_1184.dm_build_1198 == 0 {
		return
	}
	if dm_build_1183.dm_build_1143 == nil {
		dm_build_1183.dm_build_1143 = dm_build_1184
	} else {
		dm_build_1183.dm_build_1142.PushBack(dm_build_1184)
	}
	dm_build_1183.dm_build_1144 += dm_build_1184.dm_build_1198
}

func (dm_build_1186 *Dm_build_1141) dm_build_1185() {
	var dm_build_1187 = dm_build_1186.dm_build_1142.Front()
	if dm_build_1187 == nil {
		dm_build_1186.dm_build_1143 = nil
	} else {
		dm_build_1186.dm_build_1143 = dm_build_1187.Value.(*dm_build_1195)
		dm_build_1186.dm_build_1142.Remove(dm_build_1187)
	}
}

func (dm_build_1189 *Dm_build_1141) Dm_build_1188() []byte {
	var dm_build_1190 = make([]byte, dm_build_1189.dm_build_1144)
	var dm_build_1191 = dm_build_1189.dm_build_1143
	var dm_build_1192 = 0
	var dm_build_1193 = len(dm_build_1190)
	var dm_build_1194 = 0
	for dm_build_1191 != nil {
		if dm_build_1191.dm_build_1198 > 0 {
			if dm_build_1193 > dm_build_1191.dm_build_1198 {
				dm_build_1194 = dm_build_1191.dm_build_1198
			} else {
				dm_build_1194 = dm_build_1193
			}
			copy(dm_build_1190[dm_build_1192:dm_build_1192+dm_build_1194], dm_build_1191.dm_build_1196[dm_build_1191.dm_build_1197:dm_build_1191.dm_build_1197+dm_build_1194])
			dm_build_1192 += dm_build_1194
			dm_build_1193 -= dm_build_1194
		}
		if dm_build_1189.dm_build_1142.Front() == nil {
			dm_build_1191 = nil
		} else {
			dm_build_1191 = dm_build_1189.dm_build_1142.Front().Value.(*dm_build_1195)
		}
	}
	return dm_build_1190
}

type dm_build_1195 struct {
	dm_build_1196 []byte
	dm_build_1197 int
	dm_build_1198 int
}

func dm_build_1199(dm_build_1200 []byte, dm_build_1201 int, dm_build_1202 int) *dm_build_1195 {
	return &dm_build_1195{
		dm_build_1200,
		dm_build_1201,
		dm_build_1202,
	}
}

func (dm_build_1204 *dm_build_1195) dm_build_1203(dm_build_1205 *Dm_build_1219, dm_build_1206 int) int {
	if dm_build_1204.dm_build_1198 <= dm_build_1206 {
		dm_build_1206 = dm_build_1204.dm_build_1198
	}
	dm_build_1205.Dm_build_1298(dm_build_1204.dm_build_1196[dm_build_1204.dm_build_1197 : dm_build_1204.dm_build_1197+dm_build_1206])
	dm_build_1204.dm_build_1197 += dm_build_1206
	dm_build_1204.dm_build_1198 -= dm_build_1206
	return dm_build_1206
}

func (dm_build_1208 *dm_build_1195) dm_build_1207(dm_build_1209 []byte, dm_build_1210 int, dm_build_1211 int) int {
	if dm_build_1208.dm_build_1198 <= dm_build_1211 {
		dm_build_1211 = dm_build_1208.dm_build_1198
	}
	copy(dm_build_1209[dm_build_1210:dm_build_1210+dm_build_1211], dm_build_1208.dm_build_1196[dm_build_1208.dm_build_1197:dm_build_1208.dm_build_1197+dm_build_1211])
	dm_build_1208.dm_build_1197 += dm_build_1211
	dm_build_1208.dm_build_1198 -= dm_build_1211
	return dm_build_1211
}

func (dm_build_1213 *dm_build_1195) dm_build_1212(dm_build_1214 io.Writer, dm_build_1215 int) int {
	if dm_build_1213.dm_build_1198 <= dm_build_1215 {
		dm_build_1215 = dm_build_1213.dm_build_1198
	}
	dm_build_1214.Write(dm_build_1213.dm_build_1196[dm_build_1213.dm_build_1197 : dm_build_1213.dm_build_1197+dm_build_1215])
	dm_build_1213.dm_build_1197 += dm_build_1215
	dm_build_1213.dm_build_1198 -= dm_build_1215
	return dm_build_1215
}
func (dm_build_1217 *dm_build_1195) dm_build_1216(dm_build_1218 int) byte {
	return dm_build_1217.dm_build_1196[dm_build_1217.dm_build_1197+dm_build_1218]
}
