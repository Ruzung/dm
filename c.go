/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_1219 struct {
	dm_build_1220 []byte
	dm_build_1221 int
}

func Dm_build_1222(dm_build_1223 int) *Dm_build_1219 {
	return &Dm_build_1219{make([]byte, 0, dm_build_1223), 0}
}

func Dm_build_1224(dm_build_1225 []byte) *Dm_build_1219 {
	return &Dm_build_1219{dm_build_1225, 0}
}

func (dm_build_1227 *Dm_build_1219) dm_build_1226(dm_build_1228 int) *Dm_build_1219 {

	dm_build_1229 := len(dm_build_1227.dm_build_1220)
	dm_build_1230 := cap(dm_build_1227.dm_build_1220)

	if dm_build_1229+dm_build_1228 <= dm_build_1230 {
		dm_build_1227.dm_build_1220 = dm_build_1227.dm_build_1220[:dm_build_1229+dm_build_1228]
	} else {
		remain := dm_build_1228 + dm_build_1229 - dm_build_1230
		nbuf := make([]byte, dm_build_1228+dm_build_1229, 2*dm_build_1230+remain)
		copy(nbuf, dm_build_1227.dm_build_1220)
		dm_build_1227.dm_build_1220 = nbuf
	}

	return dm_build_1227
}

func (dm_build_1232 *Dm_build_1219) Dm_build_1231() int {
	return len(dm_build_1232.dm_build_1220)
}

func (dm_build_1234 *Dm_build_1219) Dm_build_1233(dm_build_1235 int) *Dm_build_1219 {
	dm_build_1234.dm_build_1220 = dm_build_1234.dm_build_1220[:dm_build_1235]
	return dm_build_1234
}

func (dm_build_1237 *Dm_build_1219) Dm_build_1236(dm_build_1238 int) *Dm_build_1219 {
	dm_build_1237.dm_build_1221 = dm_build_1238
	return dm_build_1237
}

func (dm_build_1240 *Dm_build_1219) Dm_build_1239() int {
	return dm_build_1240.dm_build_1221
}

func (dm_build_1242 *Dm_build_1219) Dm_build_1241(dm_build_1243 bool) int {
	return len(dm_build_1242.dm_build_1220) - dm_build_1242.dm_build_1221
}

func (dm_build_1245 *Dm_build_1219) Dm_build_1244(dm_build_1246 int, dm_build_1247 bool, dm_build_1248 bool) *Dm_build_1219 {

	if dm_build_1247 {
		if dm_build_1248 {
			dm_build_1245.dm_build_1226(dm_build_1246)
		} else {
			dm_build_1245.dm_build_1220 = dm_build_1245.dm_build_1220[:len(dm_build_1245.dm_build_1220)-dm_build_1246]
		}
	} else {
		if dm_build_1248 {
			dm_build_1245.dm_build_1221 += dm_build_1246
		} else {
			dm_build_1245.dm_build_1221 -= dm_build_1246
		}
	}

	return dm_build_1245
}

func (dm_build_1250 *Dm_build_1219) Dm_build_1249(dm_build_1251 io.Reader, dm_build_1252 int) int {
	dm_build_1253 := len(dm_build_1250.dm_build_1220)
	dm_build_1250.dm_build_1226(dm_build_1252)
	dm_build_1254 := 0
	for dm_build_1252 > 0 {
		n, err := dm_build_1251.Read(dm_build_1250.dm_build_1220[dm_build_1253+dm_build_1254:])
		if n > 0 && err == io.EOF {
			dm_build_1254 += n
			dm_build_1250.dm_build_1220 = dm_build_1250.dm_build_1220[:dm_build_1253+dm_build_1254]
			return dm_build_1254
		} else if n > 0 && err == nil {
			dm_build_1252 -= n
			dm_build_1254 += n
		} else if n == 0 && err != nil {
			panic("load err")
		}
	}

	return dm_build_1254
}

func (dm_build_1256 *Dm_build_1219) Dm_build_1255(dm_build_1257 io.Writer) *Dm_build_1219 {
	dm_build_1257.Write(dm_build_1256.dm_build_1220)
	return dm_build_1256
}

func (dm_build_1259 *Dm_build_1219) Dm_build_1258(dm_build_1260 bool) int {
	dm_build_1261 := len(dm_build_1259.dm_build_1220)
	dm_build_1259.dm_build_1226(1)

	if dm_build_1260 {
		return copy(dm_build_1259.dm_build_1220[dm_build_1261:], []byte{1})
	} else {
		return copy(dm_build_1259.dm_build_1220[dm_build_1261:], []byte{0})
	}
}

func (dm_build_1263 *Dm_build_1219) Dm_build_1262(dm_build_1264 byte) int {
	dm_build_1265 := len(dm_build_1263.dm_build_1220)
	dm_build_1263.dm_build_1226(1)

	return copy(dm_build_1263.dm_build_1220[dm_build_1265:], Dm_build_865.Dm_build_1040(dm_build_1264))
}

func (dm_build_1267 *Dm_build_1219) Dm_build_1266(dm_build_1268 int16) int {
	dm_build_1269 := len(dm_build_1267.dm_build_1220)
	dm_build_1267.dm_build_1226(2)

	return copy(dm_build_1267.dm_build_1220[dm_build_1269:], Dm_build_865.Dm_build_1043(dm_build_1268))
}

func (dm_build_1271 *Dm_build_1219) Dm_build_1270(dm_build_1272 int32) int {
	dm_build_1273 := len(dm_build_1271.dm_build_1220)
	dm_build_1271.dm_build_1226(4)

	return copy(dm_build_1271.dm_build_1220[dm_build_1273:], Dm_build_865.Dm_build_1046(dm_build_1272))
}

func (dm_build_1275 *Dm_build_1219) Dm_build_1274(dm_build_1276 uint8) int {
	dm_build_1277 := len(dm_build_1275.dm_build_1220)
	dm_build_1275.dm_build_1226(1)

	return copy(dm_build_1275.dm_build_1220[dm_build_1277:], Dm_build_865.Dm_build_1058(dm_build_1276))
}

func (dm_build_1279 *Dm_build_1219) Dm_build_1278(dm_build_1280 uint16) int {
	dm_build_1281 := len(dm_build_1279.dm_build_1220)
	dm_build_1279.dm_build_1226(2)

	return copy(dm_build_1279.dm_build_1220[dm_build_1281:], Dm_build_865.Dm_build_1061(dm_build_1280))
}

func (dm_build_1283 *Dm_build_1219) Dm_build_1282(dm_build_1284 uint32) int {
	dm_build_1285 := len(dm_build_1283.dm_build_1220)
	dm_build_1283.dm_build_1226(4)

	return copy(dm_build_1283.dm_build_1220[dm_build_1285:], Dm_build_865.Dm_build_1064(dm_build_1284))
}

func (dm_build_1287 *Dm_build_1219) Dm_build_1286(dm_build_1288 uint64) int {
	dm_build_1289 := len(dm_build_1287.dm_build_1220)
	dm_build_1287.dm_build_1226(8)

	return copy(dm_build_1287.dm_build_1220[dm_build_1289:], Dm_build_865.Dm_build_1067(dm_build_1288))
}

func (dm_build_1291 *Dm_build_1219) Dm_build_1290(dm_build_1292 float32) int {
	dm_build_1293 := len(dm_build_1291.dm_build_1220)
	dm_build_1291.dm_build_1226(4)

	return copy(dm_build_1291.dm_build_1220[dm_build_1293:], Dm_build_865.Dm_build_1064(math.Float32bits(dm_build_1292)))
}

func (dm_build_1295 *Dm_build_1219) Dm_build_1294(dm_build_1296 float64) int {
	dm_build_1297 := len(dm_build_1295.dm_build_1220)
	dm_build_1295.dm_build_1226(8)

	return copy(dm_build_1295.dm_build_1220[dm_build_1297:], Dm_build_865.Dm_build_1067(math.Float64bits(dm_build_1296)))
}

func (dm_build_1299 *Dm_build_1219) Dm_build_1298(dm_build_1300 []byte) int {
	dm_build_1301 := len(dm_build_1299.dm_build_1220)
	dm_build_1299.dm_build_1226(len(dm_build_1300))
	return copy(dm_build_1299.dm_build_1220[dm_build_1301:], dm_build_1300)
}

func (dm_build_1303 *Dm_build_1219) Dm_build_1302(dm_build_1304 []byte) int {
	return dm_build_1303.Dm_build_1270(int32(len(dm_build_1304))) + dm_build_1303.Dm_build_1298(dm_build_1304)
}

func (dm_build_1306 *Dm_build_1219) Dm_build_1305(dm_build_1307 []byte) int {
	return dm_build_1306.Dm_build_1274(uint8(len(dm_build_1307))) + dm_build_1306.Dm_build_1298(dm_build_1307)
}

func (dm_build_1309 *Dm_build_1219) Dm_build_1308(dm_build_1310 []byte) int {
	return dm_build_1309.Dm_build_1278(uint16(len(dm_build_1310))) + dm_build_1309.Dm_build_1298(dm_build_1310)
}

func (dm_build_1312 *Dm_build_1219) Dm_build_1311(dm_build_1313 []byte) int {
	return dm_build_1312.Dm_build_1298(dm_build_1313) + dm_build_1312.Dm_build_1262(0)
}

func (dm_build_1315 *Dm_build_1219) Dm_build_1314(dm_build_1316 string, dm_build_1317 string, dm_build_1318 *DmConnection) int {
	dm_build_1319 := Dm_build_865.Dm_build_1075(dm_build_1316, dm_build_1317, dm_build_1318)
	return dm_build_1315.Dm_build_1302(dm_build_1319)
}

func (dm_build_1321 *Dm_build_1219) Dm_build_1320(dm_build_1322 string, dm_build_1323 string, dm_build_1324 *DmConnection) int {
	dm_build_1325 := Dm_build_865.Dm_build_1075(dm_build_1322, dm_build_1323, dm_build_1324)
	return dm_build_1321.Dm_build_1305(dm_build_1325)
}

func (dm_build_1327 *Dm_build_1219) Dm_build_1326(dm_build_1328 string, dm_build_1329 string, dm_build_1330 *DmConnection) int {
	dm_build_1331 := Dm_build_865.Dm_build_1075(dm_build_1328, dm_build_1329, dm_build_1330)
	return dm_build_1327.Dm_build_1308(dm_build_1331)
}

func (dm_build_1333 *Dm_build_1219) Dm_build_1332(dm_build_1334 string, dm_build_1335 string, dm_build_1336 *DmConnection) int {
	dm_build_1337 := Dm_build_865.Dm_build_1075(dm_build_1334, dm_build_1335, dm_build_1336)
	return dm_build_1333.Dm_build_1311(dm_build_1337)
}

func (dm_build_1339 *Dm_build_1219) Dm_build_1338() byte {
	dm_build_1340 := Dm_build_865.Dm_build_958(dm_build_1339.dm_build_1220, dm_build_1339.dm_build_1221)
	dm_build_1339.dm_build_1221++
	return dm_build_1340
}

func (dm_build_1342 *Dm_build_1219) Dm_build_1341() int16 {
	dm_build_1343 := Dm_build_865.Dm_build_962(dm_build_1342.dm_build_1220, dm_build_1342.dm_build_1221)
	dm_build_1342.dm_build_1221 += 2
	return dm_build_1343
}

func (dm_build_1345 *Dm_build_1219) Dm_build_1344() int32 {
	dm_build_1346 := Dm_build_865.Dm_build_967(dm_build_1345.dm_build_1220, dm_build_1345.dm_build_1221)
	dm_build_1345.dm_build_1221 += 4
	return dm_build_1346
}

func (dm_build_1348 *Dm_build_1219) Dm_build_1347() int64 {
	dm_build_1349 := Dm_build_865.Dm_build_972(dm_build_1348.dm_build_1220, dm_build_1348.dm_build_1221)
	dm_build_1348.dm_build_1221 += 8
	return dm_build_1349
}

func (dm_build_1351 *Dm_build_1219) Dm_build_1350() float32 {
	dm_build_1352 := Dm_build_865.Dm_build_977(dm_build_1351.dm_build_1220, dm_build_1351.dm_build_1221)
	dm_build_1351.dm_build_1221 += 4
	return dm_build_1352
}

func (dm_build_1354 *Dm_build_1219) Dm_build_1353() float64 {
	dm_build_1355 := Dm_build_865.Dm_build_981(dm_build_1354.dm_build_1220, dm_build_1354.dm_build_1221)
	dm_build_1354.dm_build_1221 += 8
	return dm_build_1355
}

func (dm_build_1357 *Dm_build_1219) Dm_build_1356() uint8 {
	dm_build_1358 := Dm_build_865.Dm_build_985(dm_build_1357.dm_build_1220, dm_build_1357.dm_build_1221)
	dm_build_1357.dm_build_1221 += 1
	return dm_build_1358
}

func (dm_build_1360 *Dm_build_1219) Dm_build_1359() uint16 {
	dm_build_1361 := Dm_build_865.Dm_build_989(dm_build_1360.dm_build_1220, dm_build_1360.dm_build_1221)
	dm_build_1360.dm_build_1221 += 2
	return dm_build_1361
}

func (dm_build_1363 *Dm_build_1219) Dm_build_1362() uint32 {
	dm_build_1364 := Dm_build_865.Dm_build_994(dm_build_1363.dm_build_1220, dm_build_1363.dm_build_1221)
	dm_build_1363.dm_build_1221 += 4
	return dm_build_1364
}

func (dm_build_1366 *Dm_build_1219) Dm_build_1365(dm_build_1367 int) []byte {
	dm_build_1368 := Dm_build_865.Dm_build_1014(dm_build_1366.dm_build_1220, dm_build_1366.dm_build_1221, dm_build_1367)
	dm_build_1366.dm_build_1221 += dm_build_1367
	return dm_build_1368
}

func (dm_build_1370 *Dm_build_1219) Dm_build_1369() []byte {
	return dm_build_1370.Dm_build_1365(int(dm_build_1370.Dm_build_1344()))
}

func (dm_build_1372 *Dm_build_1219) Dm_build_1371() []byte {
	return dm_build_1372.Dm_build_1365(int(dm_build_1372.Dm_build_1338()))
}

func (dm_build_1374 *Dm_build_1219) Dm_build_1373() []byte {
	return dm_build_1374.Dm_build_1365(int(dm_build_1374.Dm_build_1341()))
}

func (dm_build_1376 *Dm_build_1219) Dm_build_1375(dm_build_1377 int) []byte {
	return dm_build_1376.Dm_build_1365(dm_build_1377)
}

func (dm_build_1379 *Dm_build_1219) Dm_build_1378() []byte {
	dm_build_1380 := 0
	for dm_build_1379.Dm_build_1338() != 0 {
		dm_build_1380++
	}
	dm_build_1379.Dm_build_1244(dm_build_1380, false, false)
	return dm_build_1379.Dm_build_1365(dm_build_1380)
}

func (dm_build_1382 *Dm_build_1219) Dm_build_1381(dm_build_1383 int, dm_build_1384 string, dm_build_1385 *DmConnection) string {
	return Dm_build_865.Dm_build_1112(dm_build_1382.Dm_build_1365(dm_build_1383), dm_build_1384, dm_build_1385)
}

func (dm_build_1387 *Dm_build_1219) Dm_build_1386(dm_build_1388 string, dm_build_1389 *DmConnection) string {
	return Dm_build_865.Dm_build_1112(dm_build_1387.Dm_build_1369(), dm_build_1388, dm_build_1389)
}

func (dm_build_1391 *Dm_build_1219) Dm_build_1390(dm_build_1392 string, dm_build_1393 *DmConnection) string {
	return Dm_build_865.Dm_build_1112(dm_build_1391.Dm_build_1371(), dm_build_1392, dm_build_1393)
}

func (dm_build_1395 *Dm_build_1219) Dm_build_1394(dm_build_1396 string, dm_build_1397 *DmConnection) string {
	return Dm_build_865.Dm_build_1112(dm_build_1395.Dm_build_1373(), dm_build_1396, dm_build_1397)
}

func (dm_build_1399 *Dm_build_1219) Dm_build_1398(dm_build_1400 string, dm_build_1401 *DmConnection) string {
	return Dm_build_865.Dm_build_1112(dm_build_1399.Dm_build_1378(), dm_build_1400, dm_build_1401)
}

func (dm_build_1403 *Dm_build_1219) Dm_build_1402(dm_build_1404 int, dm_build_1405 byte) int {
	return dm_build_1403.Dm_build_1438(dm_build_1404, Dm_build_865.Dm_build_1040(dm_build_1405))
}

func (dm_build_1407 *Dm_build_1219) Dm_build_1406(dm_build_1408 int, dm_build_1409 int16) int {
	return dm_build_1407.Dm_build_1438(dm_build_1408, Dm_build_865.Dm_build_1043(dm_build_1409))
}

func (dm_build_1411 *Dm_build_1219) Dm_build_1410(dm_build_1412 int, dm_build_1413 int32) int {
	return dm_build_1411.Dm_build_1438(dm_build_1412, Dm_build_865.Dm_build_1046(dm_build_1413))
}

func (dm_build_1415 *Dm_build_1219) Dm_build_1414(dm_build_1416 int, dm_build_1417 int64) int {
	return dm_build_1415.Dm_build_1438(dm_build_1416, Dm_build_865.Dm_build_1049(dm_build_1417))
}

func (dm_build_1419 *Dm_build_1219) Dm_build_1418(dm_build_1420 int, dm_build_1421 float32) int {
	return dm_build_1419.Dm_build_1438(dm_build_1420, Dm_build_865.Dm_build_1052(dm_build_1421))
}

func (dm_build_1423 *Dm_build_1219) Dm_build_1422(dm_build_1424 int, dm_build_1425 float64) int {
	return dm_build_1423.Dm_build_1438(dm_build_1424, Dm_build_865.Dm_build_1055(dm_build_1425))
}

func (dm_build_1427 *Dm_build_1219) Dm_build_1426(dm_build_1428 int, dm_build_1429 uint8) int {
	return dm_build_1427.Dm_build_1438(dm_build_1428, Dm_build_865.Dm_build_1058(dm_build_1429))
}

func (dm_build_1431 *Dm_build_1219) Dm_build_1430(dm_build_1432 int, dm_build_1433 uint16) int {
	return dm_build_1431.Dm_build_1438(dm_build_1432, Dm_build_865.Dm_build_1061(dm_build_1433))
}

func (dm_build_1435 *Dm_build_1219) Dm_build_1434(dm_build_1436 int, dm_build_1437 uint32) int {
	return dm_build_1435.Dm_build_1438(dm_build_1436, Dm_build_865.Dm_build_1064(dm_build_1437))
}

func (dm_build_1439 *Dm_build_1219) Dm_build_1438(dm_build_1440 int, dm_build_1441 []byte) int {
	return copy(dm_build_1439.dm_build_1220[dm_build_1440:], dm_build_1441)
}

func (dm_build_1443 *Dm_build_1219) Dm_build_1442(dm_build_1444 int, dm_build_1445 []byte) int {
	return dm_build_1443.Dm_build_1410(dm_build_1444, int32(len(dm_build_1445))) + dm_build_1443.Dm_build_1438(dm_build_1444+4, dm_build_1445)
}

func (dm_build_1447 *Dm_build_1219) Dm_build_1446(dm_build_1448 int, dm_build_1449 []byte) int {
	return dm_build_1447.Dm_build_1402(dm_build_1448, byte(len(dm_build_1449))) + dm_build_1447.Dm_build_1438(dm_build_1448+1, dm_build_1449)
}

func (dm_build_1451 *Dm_build_1219) Dm_build_1450(dm_build_1452 int, dm_build_1453 []byte) int {
	return dm_build_1451.Dm_build_1406(dm_build_1452, int16(len(dm_build_1453))) + dm_build_1451.Dm_build_1438(dm_build_1452+2, dm_build_1453)
}

func (dm_build_1455 *Dm_build_1219) Dm_build_1454(dm_build_1456 int, dm_build_1457 []byte) int {
	return dm_build_1455.Dm_build_1438(dm_build_1456, dm_build_1457) + dm_build_1455.Dm_build_1402(dm_build_1456+len(dm_build_1457), 0)
}

func (dm_build_1459 *Dm_build_1219) Dm_build_1458(dm_build_1460 int, dm_build_1461 string, dm_build_1462 string, dm_build_1463 *DmConnection) int {
	return dm_build_1459.Dm_build_1442(dm_build_1460, Dm_build_865.Dm_build_1075(dm_build_1461, dm_build_1462, dm_build_1463))
}

func (dm_build_1465 *Dm_build_1219) Dm_build_1464(dm_build_1466 int, dm_build_1467 string, dm_build_1468 string, dm_build_1469 *DmConnection) int {
	return dm_build_1465.Dm_build_1446(dm_build_1466, Dm_build_865.Dm_build_1075(dm_build_1467, dm_build_1468, dm_build_1469))
}

func (dm_build_1471 *Dm_build_1219) Dm_build_1470(dm_build_1472 int, dm_build_1473 string, dm_build_1474 string, dm_build_1475 *DmConnection) int {
	return dm_build_1471.Dm_build_1450(dm_build_1472, Dm_build_865.Dm_build_1075(dm_build_1473, dm_build_1474, dm_build_1475))
}

func (dm_build_1477 *Dm_build_1219) Dm_build_1476(dm_build_1478 int, dm_build_1479 string, dm_build_1480 string, dm_build_1481 *DmConnection) int {
	return dm_build_1477.Dm_build_1454(dm_build_1478, Dm_build_865.Dm_build_1075(dm_build_1479, dm_build_1480, dm_build_1481))
}

func (dm_build_1483 *Dm_build_1219) Dm_build_1482(dm_build_1484 int) byte {
	return Dm_build_865.Dm_build_1080(dm_build_1483.Dm_build_1509(dm_build_1484, 1))
}

func (dm_build_1486 *Dm_build_1219) Dm_build_1485(dm_build_1487 int) int16 {
	return Dm_build_865.Dm_build_1083(dm_build_1486.Dm_build_1509(dm_build_1487, 2))
}

func (dm_build_1489 *Dm_build_1219) Dm_build_1488(dm_build_1490 int) int32 {
	return Dm_build_865.Dm_build_1086(dm_build_1489.Dm_build_1509(dm_build_1490, 4))
}

func (dm_build_1492 *Dm_build_1219) Dm_build_1491(dm_build_1493 int) int64 {
	return Dm_build_865.Dm_build_1089(dm_build_1492.Dm_build_1509(dm_build_1493, 8))
}

func (dm_build_1495 *Dm_build_1219) Dm_build_1494(dm_build_1496 int) float32 {
	return Dm_build_865.Dm_build_1092(dm_build_1495.Dm_build_1509(dm_build_1496, 4))
}

func (dm_build_1498 *Dm_build_1219) Dm_build_1497(dm_build_1499 int) float64 {
	return Dm_build_865.Dm_build_1095(dm_build_1498.Dm_build_1509(dm_build_1499, 8))
}

func (dm_build_1501 *Dm_build_1219) Dm_build_1500(dm_build_1502 int) uint8 {
	return Dm_build_865.Dm_build_1098(dm_build_1501.Dm_build_1509(dm_build_1502, 1))
}

func (dm_build_1504 *Dm_build_1219) Dm_build_1503(dm_build_1505 int) uint16 {
	return Dm_build_865.Dm_build_1101(dm_build_1504.Dm_build_1509(dm_build_1505, 2))
}

func (dm_build_1507 *Dm_build_1219) Dm_build_1506(dm_build_1508 int) uint32 {
	return Dm_build_865.Dm_build_1104(dm_build_1507.Dm_build_1509(dm_build_1508, 4))
}

func (dm_build_1510 *Dm_build_1219) Dm_build_1509(dm_build_1511 int, dm_build_1512 int) []byte {
	return dm_build_1510.dm_build_1220[dm_build_1511 : dm_build_1511+dm_build_1512]
}

func (dm_build_1514 *Dm_build_1219) Dm_build_1513(dm_build_1515 int) []byte {
	dm_build_1516 := dm_build_1514.Dm_build_1488(dm_build_1515)
	return dm_build_1514.Dm_build_1509(dm_build_1515+4, int(dm_build_1516))
}

func (dm_build_1518 *Dm_build_1219) Dm_build_1517(dm_build_1519 int) []byte {
	dm_build_1520 := dm_build_1518.Dm_build_1482(dm_build_1519)
	return dm_build_1518.Dm_build_1509(dm_build_1519+1, int(dm_build_1520))
}

func (dm_build_1522 *Dm_build_1219) Dm_build_1521(dm_build_1523 int) []byte {
	dm_build_1524 := dm_build_1522.Dm_build_1485(dm_build_1523)
	return dm_build_1522.Dm_build_1509(dm_build_1523+2, int(dm_build_1524))
}

func (dm_build_1526 *Dm_build_1219) Dm_build_1525(dm_build_1527 int) []byte {
	dm_build_1528 := 0
	for dm_build_1526.Dm_build_1482(dm_build_1527) != 0 {
		dm_build_1527++
		dm_build_1528++
	}

	return dm_build_1526.Dm_build_1509(dm_build_1527-dm_build_1528, int(dm_build_1528))
}

func (dm_build_1530 *Dm_build_1219) Dm_build_1529(dm_build_1531 int, dm_build_1532 string, dm_build_1533 *DmConnection) string {
	return Dm_build_865.Dm_build_1112(dm_build_1530.Dm_build_1513(dm_build_1531), dm_build_1532, dm_build_1533)
}

func (dm_build_1535 *Dm_build_1219) Dm_build_1534(dm_build_1536 int, dm_build_1537 string, dm_build_1538 *DmConnection) string {
	return Dm_build_865.Dm_build_1112(dm_build_1535.Dm_build_1517(dm_build_1536), dm_build_1537, dm_build_1538)
}

func (dm_build_1540 *Dm_build_1219) Dm_build_1539(dm_build_1541 int, dm_build_1542 string, dm_build_1543 *DmConnection) string {
	return Dm_build_865.Dm_build_1112(dm_build_1540.Dm_build_1521(dm_build_1541), dm_build_1542, dm_build_1543)
}

func (dm_build_1545 *Dm_build_1219) Dm_build_1544(dm_build_1546 int, dm_build_1547 string, dm_build_1548 *DmConnection) string {
	return Dm_build_865.Dm_build_1112(dm_build_1545.Dm_build_1525(dm_build_1546), dm_build_1547, dm_build_1548)
}
