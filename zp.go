/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"dm/util"
	"os"
	"strconv"
	"strings"
)

const (
	Dm_build_266 = "7.6.0.0"

	Dm_build_267 = "7.0.0.9"

	Dm_build_268 = "8.0.0.73"

	Dm_build_269 = "7.1.2.128"

	Dm_build_270 = "7.1.5.144"

	Dm_build_271 = "7.1.6.123"

	Dm_build_272 = 32768 - 128

	Dm_build_273 int16 = 1

	Dm_build_274 int16 = 2

	Dm_build_275 int16 = 3

	Dm_build_276 int16 = 4

	Dm_build_277 int16 = 5

	Dm_build_278 int16 = 6

	Dm_build_279 int16 = 7

	Dm_build_280 int16 = 8

	Dm_build_281 int16 = 9

	Dm_build_282 int16 = 13

	Dm_build_283 int16 = 14

	Dm_build_284 int16 = 15

	Dm_build_285 int16 = 17

	Dm_build_286 int16 = 21

	Dm_build_287 int16 = 24

	Dm_build_288 int16 = 27

	Dm_build_289 int16 = 29

	Dm_build_290 int16 = 30

	Dm_build_291 int16 = 31

	Dm_build_292 int16 = 32

	Dm_build_293 int16 = 44

	Dm_build_294 int16 = 52

	Dm_build_295 int16 = 60

	Dm_build_296 int16 = 71

	Dm_build_297 int16 = 90

	Dm_build_298 int16 = 91

	Dm_build_299 int16 = 200

	Dm_build_300 = 64

	Dm_build_301 = 20

	Dm_build_302 = 0

	Dm_build_303 = 4

	Dm_build_304 = 6

	Dm_build_305 = 10

	Dm_build_306 = 14

	Dm_build_307 = 18

	Dm_build_308 = 19

	Dm_build_309 = 128

	Dm_build_310 = 256

	Dm_build_311 = 0xffff

	Dm_build_312 int32 = 2

	Dm_build_313 = -1

	Dm_build_314 uint16 = 0xFFFE

	Dm_build_315 uint16 = uint16(Dm_build_314 - 3)

	Dm_build_316 uint16 = Dm_build_314

	Dm_build_317 int32 = 0xFFFF

	Dm_build_318 int32 = 0x80

	Dm_build_319 byte = 0x60

	Dm_build_320 uint16 = uint16(Dm_build_316)

	Dm_build_321 uint16 = uint16(Dm_build_317)

	Dm_build_322 int16 = 0x00

	Dm_build_323 int16 = 0x03

	Dm_build_324 int32 = 0x80

	Dm_build_325 byte = 0

	Dm_build_326 byte = 1

	Dm_build_327 byte = 2

	Dm_build_328 byte = 3

	Dm_build_329 byte = 4

	Dm_build_330 byte = Dm_build_325

	Dm_build_331 int = 10

	Dm_build_332 int32 = 32

	Dm_build_333 int32 = 65536

	Dm_build_334 byte = 0

	Dm_build_335 byte = 1

	Dm_build_336 int32 = 0x00000000

	Dm_build_337 int32 = 0x00000020

	Dm_build_338 int32 = 0x00000040

	Dm_build_339 int32 = 0x00000FFF

	Dm_build_340 int32 = 0

	Dm_build_341 int32 = 1

	Dm_build_342 int32 = 2

	Dm_build_343 int32 = 3

	Dm_build_344 = 8192

	Dm_build_345 = 1

	Dm_build_346 = 2

	Dm_build_347 = 0

	Dm_build_348 = 0

	Dm_build_349 = 1

	Dm_build_350 = -1

	Dm_build_351 int16 = 0

	Dm_build_352 int16 = 1

	Dm_build_353 int16 = 2

	Dm_build_354 int16 = 3

	Dm_build_355 int16 = 4

	Dm_build_356 int16 = 127

	Dm_build_357 int16 = Dm_build_356 + 20

	Dm_build_358 int16 = Dm_build_356 + 21

	Dm_build_359 int16 = Dm_build_356 + 22

	Dm_build_360 int16 = Dm_build_356 + 24

	Dm_build_361 int16 = Dm_build_356 + 25

	Dm_build_362 int16 = Dm_build_356 + 26

	Dm_build_363 int16 = Dm_build_356 + 30

	Dm_build_364 int16 = Dm_build_356 + 31

	Dm_build_365 int16 = Dm_build_356 + 32

	Dm_build_366 int16 = Dm_build_356 + 33

	Dm_build_367 int16 = Dm_build_356 + 35

	Dm_build_368 int16 = Dm_build_356 + 38

	Dm_build_369 int16 = Dm_build_356 + 39

	Dm_build_370 int16 = Dm_build_356 + 51

	Dm_build_371 int16 = Dm_build_356 + 71

	Dm_build_372 int16 = Dm_build_356 + 124

	Dm_build_373 int16 = Dm_build_356 + 125

	Dm_build_374 int16 = Dm_build_356 + 126

	Dm_build_375 int16 = Dm_build_356 + 127

	Dm_build_376 int16 = Dm_build_356 + 128

	Dm_build_377 int16 = Dm_build_356 + 129

	Dm_build_378 byte = 0

	Dm_build_379 byte = 2

	Dm_build_380 = 2048

	Dm_build_381 = -1

	Dm_build_382 = 0

	Dm_build_383 = 16000

	Dm_build_384 = 32000

	Dm_build_385 = 0x00000000

	Dm_build_386 = 0x00000020

	Dm_build_387 = 0x00000040

	Dm_build_388 = 0x00000FFF
)

type dm_build_389 interface {
	dm_build_390()
	dm_build_391() error
	dm_build_392()
	dm_build_393(imsg dm_build_389) error
	dm_build_394() error
	dm_build_395() (interface{}, error)
	dm_build_396()
	dm_build_397(imsg dm_build_389) (interface{}, error)
	dm_build_398()
	dm_build_399() error
	dm_build_400() byte
	dm_build_401() int32
	dm_build_402(length int32)
	dm_build_403() int16
}

type dm_build_404 struct {
	dm_build_405 *dm_build_2

	dm_build_406 int16

	dm_build_407 int32

	dm_build_408 *DmStatement
}

func (dm_build_410 *dm_build_404) dm_build_409(dm_build_411 *dm_build_2, dm_build_412 int16) *dm_build_404 {
	dm_build_410.dm_build_405 = dm_build_411
	dm_build_410.dm_build_406 = dm_build_412
	return dm_build_410
}

func (dm_build_414 *dm_build_404) dm_build_413(dm_build_415 *dm_build_2, dm_build_416 int16, dm_build_417 *DmStatement) *dm_build_404 {
	dm_build_414.dm_build_409(dm_build_415, dm_build_416).dm_build_408 = dm_build_417
	return dm_build_414
}

func dm_build_418(dm_build_419 *dm_build_2, dm_build_420 int16) *dm_build_404 {
	return new(dm_build_404).dm_build_409(dm_build_419, dm_build_420)
}

func dm_build_421(dm_build_422 *dm_build_2, dm_build_423 int16, dm_build_424 *DmStatement) *dm_build_404 {
	return new(dm_build_404).dm_build_413(dm_build_422, dm_build_423, dm_build_424)
}

func (dm_build_426 *dm_build_404) dm_build_390() {
	dm_build_426.dm_build_405.dm_build_5.Dm_build_1233(0)
	dm_build_426.dm_build_405.dm_build_5.Dm_build_1244(Dm_build_300, true, true)
}

func (dm_build_428 *dm_build_404) dm_build_391() error {
	return nil
}

func (dm_build_430 *dm_build_404) dm_build_392() {
	if dm_build_430.dm_build_408 == nil {
		dm_build_430.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_302, 0)
	} else {
		dm_build_430.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_302, dm_build_430.dm_build_408.id)
	}

	dm_build_430.dm_build_405.dm_build_5.Dm_build_1406(Dm_build_303, dm_build_430.dm_build_406)
	dm_build_430.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_304, int32(dm_build_430.dm_build_405.dm_build_5.Dm_build_1231()-Dm_build_300))
}

func (dm_build_432 *dm_build_404) dm_build_394() error {
	dm_build_432.dm_build_405.dm_build_5.Dm_build_1236(0)
	dm_build_432.dm_build_405.dm_build_5.Dm_build_1244(Dm_build_300, false, true)
	return dm_build_432.dm_build_437()
}

func (dm_build_434 *dm_build_404) dm_build_395() (interface{}, error) {
	return nil, nil
}

func (dm_build_436 *dm_build_404) dm_build_396() {
}

func (dm_build_438 *dm_build_404) dm_build_437() error {
	dm_build_438.dm_build_407 = dm_build_438.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_305)
	if dm_build_438.dm_build_407 < 0 && dm_build_438.dm_build_407 != EC_RN_EXCEED_ROWSET_SIZE.ErrCode {
		return (&DmError{dm_build_438.dm_build_407, dm_build_438.dm_build_439(), nil, ""}).throw()
	} else if dm_build_438.dm_build_407 > 0 {

	} else if dm_build_438.dm_build_406 == Dm_build_299 || dm_build_438.dm_build_406 == Dm_build_273 {
		dm_build_438.dm_build_439()
	}

	return nil
}

func (dm_build_440 *dm_build_404) dm_build_439() string {

	dm_build_441 := dm_build_440.dm_build_405.dm_build_6.getServerEncoding()

	if dm_build_441 != "" && dm_build_441 == ENCODING_EUCKR && Locale != LANGUAGE_EN {
		dm_build_441 = ENCODING_GB18030
	}

	dm_build_440.dm_build_405.dm_build_5.Dm_build_1244(int(dm_build_440.dm_build_405.dm_build_5.Dm_build_1344()), false, true)

	dm_build_440.dm_build_405.dm_build_5.Dm_build_1244(int(dm_build_440.dm_build_405.dm_build_5.Dm_build_1344()), false, true)

	dm_build_440.dm_build_405.dm_build_5.Dm_build_1244(int(dm_build_440.dm_build_405.dm_build_5.Dm_build_1344()), false, true)

	return dm_build_440.dm_build_405.dm_build_5.Dm_build_1386(dm_build_441, dm_build_440.dm_build_405.dm_build_6)
}

func (dm_build_443 *dm_build_404) dm_build_393(dm_build_444 dm_build_389) (dm_build_445 error) {
	dm_build_444.dm_build_390()
	if dm_build_445 = dm_build_444.dm_build_391(); dm_build_445 != nil {
		return dm_build_445
	}
	dm_build_444.dm_build_392()
	return nil
}

func (dm_build_447 *dm_build_404) dm_build_397(dm_build_448 dm_build_389) (dm_build_449 interface{}, dm_build_450 error) {
	dm_build_450 = dm_build_448.dm_build_394()
	if dm_build_450 != nil {
		return nil, dm_build_450
	}
	dm_build_449, dm_build_450 = dm_build_448.dm_build_395()
	if dm_build_450 != nil {
		return nil, dm_build_450
	}
	dm_build_448.dm_build_396()
	return dm_build_449, nil
}

func (dm_build_452 *dm_build_404) dm_build_398() {
	dm_build_452.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_308, dm_build_452.dm_build_400())
}

func (dm_build_454 *dm_build_404) dm_build_399() error {
	dm_build_455 := dm_build_454.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_308)
	dm_build_456 := dm_build_454.dm_build_400()
	if dm_build_455 != dm_build_456 {
		return ECGO_MSG_CHECK_ERROR.throw()
	}
	return nil
}

func (dm_build_458 *dm_build_404) dm_build_400() byte {
	dm_build_459 := dm_build_458.dm_build_405.dm_build_5.Dm_build_1482(0)

	for i := 1; i < Dm_build_308; i++ {
		dm_build_459 ^= dm_build_458.dm_build_405.dm_build_5.Dm_build_1482(i)
	}

	return dm_build_459
}

func (dm_build_461 *dm_build_404) dm_build_401() int32 {
	return dm_build_461.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_304)
}

func (dm_build_463 *dm_build_404) dm_build_402(dm_build_464 int32) {
	dm_build_463.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_304, dm_build_464)
}

func (dm_build_466 *dm_build_404) dm_build_403() int16 {
	return dm_build_466.dm_build_406
}

type dm_build_467 struct {
	dm_build_404
}

func dm_build_468(dm_build_469 *dm_build_2) *dm_build_467 {
	dm_build_470 := new(dm_build_467)
	dm_build_470.dm_build_409(dm_build_469, Dm_build_280)
	return dm_build_470
}

type dm_build_471 struct {
	dm_build_404
	dm_build_472 string
}

func dm_build_473(dm_build_474 *dm_build_2, dm_build_475 *DmStatement, dm_build_476 string) *dm_build_471 {
	dm_build_477 := new(dm_build_471)
	dm_build_477.dm_build_413(dm_build_474, Dm_build_288, dm_build_475)
	dm_build_477.dm_build_472 = dm_build_476
	dm_build_477.dm_build_408.cursorName = dm_build_476
	return dm_build_477
}

func (dm_build_479 *dm_build_471) dm_build_391() error {
	dm_build_479.dm_build_405.dm_build_5.Dm_build_1332(dm_build_479.dm_build_472, dm_build_479.dm_build_405.dm_build_6.getServerEncoding(), dm_build_479.dm_build_405.dm_build_6)
	dm_build_479.dm_build_405.dm_build_5.Dm_build_1270(1)
	return nil
}

type Dm_build_480 struct {
	dm_build_496
	dm_build_481 []OptParameter
}

func dm_build_482(dm_build_483 *dm_build_2, dm_build_484 *DmStatement, dm_build_485 []OptParameter) *Dm_build_480 {
	dm_build_486 := new(Dm_build_480)
	dm_build_486.dm_build_413(dm_build_483, Dm_build_298, dm_build_484)
	dm_build_486.dm_build_481 = dm_build_485
	return dm_build_486
}

func (dm_build_488 *Dm_build_480) dm_build_391() error {
	dm_build_489 := len(dm_build_488.dm_build_481)

	dm_build_488.dm_build_510(int16(dm_build_489), 1)

	dm_build_488.dm_build_405.dm_build_5.Dm_build_1332(dm_build_488.dm_build_408.nativeSql, dm_build_488.dm_build_408.dmConn.getServerEncoding(), dm_build_488.dm_build_408.dmConn)

	for _, param := range dm_build_488.dm_build_481 {
		dm_build_488.dm_build_405.dm_build_5.Dm_build_1262(param.ioType)
		dm_build_488.dm_build_405.dm_build_5.Dm_build_1270(int32(param.tp))
		dm_build_488.dm_build_405.dm_build_5.Dm_build_1270(int32(param.prec))
		dm_build_488.dm_build_405.dm_build_5.Dm_build_1270(int32(param.scale))
	}

	for _, param := range dm_build_488.dm_build_481 {
		if param.bytes == nil {
			dm_build_488.dm_build_405.dm_build_5.Dm_build_1278(Dm_build_316)
		} else {
			dm_build_488.dm_build_405.dm_build_5.Dm_build_1308(param.bytes[:len(param.bytes)])
		}
	}
	return nil
}

func (dm_build_491 *Dm_build_480) dm_build_395() (interface{}, error) {
	return dm_build_491.dm_build_496.dm_build_395()
}

const (
	Dm_build_492 int = 0x01

	Dm_build_493 int = 0x02

	Dm_build_494 int = 0x04

	Dm_build_495 int = 0x08
)

type dm_build_496 struct {
	dm_build_404
	dm_build_497 [][]interface{}
	dm_build_498 []parameter
	dm_build_499 bool
}

func dm_build_500(dm_build_501 *dm_build_2, dm_build_502 int16, dm_build_503 *DmStatement) *dm_build_496 {
	dm_build_504 := new(dm_build_496)
	dm_build_504.dm_build_413(dm_build_501, dm_build_502, dm_build_503)
	dm_build_504.dm_build_499 = true
	return dm_build_504
}

func dm_build_505(dm_build_506 *dm_build_2, dm_build_507 *DmStatement, dm_build_508 [][]interface{}) *dm_build_496 {
	dm_build_509 := new(dm_build_496)

	if dm_build_506.dm_build_6.Execute2 {
		dm_build_509.dm_build_413(dm_build_506, Dm_build_282, dm_build_507)
	} else {
		dm_build_509.dm_build_413(dm_build_506, Dm_build_278, dm_build_507)
	}

	dm_build_509.dm_build_498 = dm_build_507.params
	dm_build_509.dm_build_497 = dm_build_508
	dm_build_509.dm_build_499 = true
	return dm_build_509
}

func (dm_build_511 *dm_build_496) dm_build_510(dm_build_512 int16, dm_build_513 int64) {

	dm_build_514 := Dm_build_301

	if dm_build_511.dm_build_405.dm_build_6.autoCommit {
		dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1402(dm_build_514, 1)
	} else {
		dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1402(dm_build_514, 0)
	}

	dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1406(dm_build_514, dm_build_512)

	dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1402(dm_build_514, 1)

	dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1414(dm_build_514, dm_build_513)

	dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1414(dm_build_514, dm_build_511.dm_build_408.cursorUpdateRow)

	if dm_build_511.dm_build_408.maxRows <= 0 || dm_build_511.dm_build_408.dmConn.dmConnector.enRsCache {
		dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1414(dm_build_514, INT64_MAX)
	} else {
		dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1414(dm_build_514, dm_build_511.dm_build_408.maxRows)
	}

	dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1402(dm_build_514, 1)

	if dm_build_511.dm_build_405.dm_build_6.dmConnector.continueBatchOnError {
		dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1402(dm_build_514, 1)
	} else {
		dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1402(dm_build_514, 0)
	}

	dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1402(dm_build_514, 0)

	dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1402(dm_build_514, 0)

	if dm_build_511.dm_build_408.queryTimeout == 0 {
		dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1410(dm_build_514, -1)
	} else {
		dm_build_514 += dm_build_511.dm_build_405.dm_build_5.Dm_build_1410(dm_build_514, dm_build_511.dm_build_408.queryTimeout)
	}
}

func (dm_build_516 *dm_build_496) dm_build_391() error {
	var dm_build_517 int16
	var dm_build_518 int64

	if dm_build_516.dm_build_498 != nil {
		dm_build_517 = int16(len(dm_build_516.dm_build_498))
	} else {
		dm_build_517 = 0
	}

	if dm_build_516.dm_build_497 != nil {
		dm_build_518 = int64(len(dm_build_516.dm_build_497))
	} else {
		dm_build_518 = 0
	}

	dm_build_516.dm_build_510(dm_build_517, dm_build_518)

	if dm_build_517 > 0 {
		err := dm_build_516.dm_build_519(dm_build_516.dm_build_498)
		if err != nil {
			return err
		}
		if dm_build_516.dm_build_497 != nil && len(dm_build_516.dm_build_497) > 0 {
			for _, paramObject := range dm_build_516.dm_build_497 {
				if err := dm_build_516.dm_build_522(paramObject); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (dm_build_520 *dm_build_496) dm_build_519(dm_build_521 []parameter) error {
	for _, param := range dm_build_521 {

		if param.colType == CURSOR && param.ioType == IO_TYPE_OUT {
			dm_build_520.dm_build_405.dm_build_5.Dm_build_1262(IO_TYPE_INOUT)
		} else {
			dm_build_520.dm_build_405.dm_build_5.Dm_build_1262(param.ioType)
		}

		dm_build_520.dm_build_405.dm_build_5.Dm_build_1270(param.colType)

		lprec := param.prec
		lscale := param.scale
		typeDesc := param.typeDescriptor
		switch param.colType {
		case ARRAY, SARRAY:
			tmp, err := getPackArraySize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case PLTYPE_RECORD:
			tmp, err := getPackRecordSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case CLASS:
			tmp, err := getPackClassSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case BLOB:
			if isComplexType(int(param.colType), int(param.scale)) {
				lprec = int32(typeDesc.getObjId())
				if lprec == 4 {
					lprec = int32(typeDesc.getOuterId())
				}
			}
		}

		dm_build_520.dm_build_405.dm_build_5.Dm_build_1270(lprec)

		dm_build_520.dm_build_405.dm_build_5.Dm_build_1270(lscale)

		switch param.colType {
		case ARRAY, SARRAY:
			err := packArray(typeDesc, dm_build_520.dm_build_405.dm_build_5)
			if err != nil {
				return err
			}

		case PLTYPE_RECORD:
			err := packRecord(typeDesc, dm_build_520.dm_build_405.dm_build_5)
			if err != nil {
				return err
			}

		case CLASS:
			err := packClass(typeDesc, dm_build_520.dm_build_405.dm_build_5)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (dm_build_523 *dm_build_496) dm_build_522(dm_build_524 []interface{}) error {
	for i := 0; i < len(dm_build_523.dm_build_498); i++ {

		if dm_build_523.dm_build_498[i].colType == CURSOR {
			dm_build_523.dm_build_405.dm_build_5.Dm_build_1266(ULINT_SIZE)
			dm_build_523.dm_build_405.dm_build_5.Dm_build_1270(dm_build_523.dm_build_498[i].cursorStmt.id)
			continue
		}

		if dm_build_523.dm_build_498[i].ioType == IO_TYPE_OUT {
			continue
		}

		switch dm_build_524[i].(type) {
		case []byte:
			if dataBytes, ok := dm_build_524[i].([]byte); ok {
				if len(dataBytes) > Dm_build_311 {
					return ECGO_DATA_TOO_LONG.throw()
				}
				dm_build_523.dm_build_405.dm_build_5.Dm_build_1308(dataBytes)
			}
		case int:
			if dm_build_524[i] == ParamDataEnum_Null {
				dm_build_523.dm_build_405.dm_build_5.Dm_build_1278(Dm_build_316)
			} else if dm_build_524[i] == ParamDataEnum_OFF_ROW {
				dm_build_523.dm_build_405.dm_build_5.Dm_build_1266(0)
			}
		case lobCtl:
			dm_build_523.dm_build_405.dm_build_5.Dm_build_1278(uint16(Dm_build_315))
			dm_build_523.dm_build_405.dm_build_5.Dm_build_1298(dm_build_524[i].(lobCtl).value)
		default:
			panic("Bind param data failed by invalid param data type: ")
		}
	}

	return nil
}

func (dm_build_526 *dm_build_496) dm_build_395() (interface{}, error) {
	dm_build_527 := execInfo{}
	dm_build_528 := dm_build_526.dm_build_408.dmConn

	dm_build_529 := Dm_build_301

	dm_build_527.retSqlType = dm_build_526.dm_build_405.dm_build_5.Dm_build_1485(dm_build_529)
	dm_build_529 += USINT_SIZE

	dm_build_530 := dm_build_526.dm_build_405.dm_build_5.Dm_build_1485(dm_build_529)
	dm_build_529 += USINT_SIZE

	dm_build_527.updateCount = dm_build_526.dm_build_405.dm_build_5.Dm_build_1491(dm_build_529)
	dm_build_529 += DDWORD_SIZE

	dm_build_531 := dm_build_526.dm_build_405.dm_build_5.Dm_build_1485(dm_build_529)
	dm_build_529 += USINT_SIZE

	dm_build_527.rsUpdatable = dm_build_526.dm_build_405.dm_build_5.Dm_build_1482(dm_build_529) != 0
	dm_build_529 += BYTE_SIZE

	dm_build_532 := dm_build_526.dm_build_405.dm_build_5.Dm_build_1485(dm_build_529)
	dm_build_529 += ULINT_SIZE

	dm_build_527.printLen = dm_build_526.dm_build_405.dm_build_5.Dm_build_1488(dm_build_529)
	dm_build_529 += ULINT_SIZE

	var dm_build_533 int16 = -1
	if dm_build_527.retSqlType == Dm_build_366 || dm_build_527.retSqlType == Dm_build_367 {
		dm_build_527.rowid = 0

		dm_build_527.rsBdta = dm_build_526.dm_build_405.dm_build_5.Dm_build_1482(dm_build_529) == Dm_build_379
		dm_build_529 += BYTE_SIZE

		dm_build_533 = dm_build_526.dm_build_405.dm_build_5.Dm_build_1485(dm_build_529)
		dm_build_529 += USINT_SIZE
		dm_build_529 += 5
	} else {
		dm_build_527.rowid = dm_build_526.dm_build_405.dm_build_5.Dm_build_1491(dm_build_529)
		dm_build_529 += DDWORD_SIZE
	}

	dm_build_527.execId = dm_build_526.dm_build_405.dm_build_5.Dm_build_1488(dm_build_529)
	dm_build_529 += ULINT_SIZE

	dm_build_527.rsCacheOffset = dm_build_526.dm_build_405.dm_build_5.Dm_build_1488(dm_build_529)
	dm_build_529 += ULINT_SIZE

	dm_build_534 := dm_build_526.dm_build_405.dm_build_5.Dm_build_1482(dm_build_529)
	dm_build_529 += BYTE_SIZE
	dm_build_535 := (dm_build_534 & 0x01) == 0x01
	dm_build_536 := (dm_build_534 & 0x02) == 0x02

	dm_build_528.TrxStatus = dm_build_526.dm_build_405.dm_build_5.Dm_build_1488(dm_build_529)
	dm_build_528.setTrxFinish(dm_build_528.TrxStatus)
	dm_build_529 += ULINT_SIZE

	if dm_build_527.printLen > 0 {
		bytes := dm_build_526.dm_build_405.dm_build_5.Dm_build_1365(int(dm_build_527.printLen))
		dm_build_527.printMsg = Dm_build_865.Dm_build_1019(bytes, 0, len(bytes), dm_build_528.getServerEncoding(), dm_build_528)
	}

	if dm_build_531 > 0 {
		dm_build_527.outParamDatas = dm_build_526.dm_build_537(int(dm_build_531))
	}

	switch dm_build_527.retSqlType {
	case Dm_build_368:
		dm_build_528.dmConnector.localTimezone = dm_build_526.dm_build_405.dm_build_5.Dm_build_1341()
	case Dm_build_366:
		dm_build_527.hasResultSet = true
		if dm_build_530 > 0 {
			dm_build_526.dm_build_408.columns = dm_build_526.dm_build_546(int(dm_build_530), dm_build_527.rsBdta)
		}
		dm_build_526.dm_build_556(&dm_build_527, len(dm_build_526.dm_build_408.columns), int(dm_build_532), int(dm_build_533))
	case Dm_build_367:
		if dm_build_530 > 0 || dm_build_532 > 0 {
			dm_build_527.hasResultSet = true
		}
		if dm_build_530 > 0 {
			dm_build_526.dm_build_408.columns = dm_build_526.dm_build_546(int(dm_build_530), dm_build_527.rsBdta)
		}
		dm_build_526.dm_build_556(&dm_build_527, len(dm_build_526.dm_build_408.columns), int(dm_build_532), int(dm_build_533))
	case Dm_build_369:
		dm_build_528.IsoLevel = int32(dm_build_526.dm_build_405.dm_build_5.Dm_build_1341())
		dm_build_528.ReadOnly = dm_build_526.dm_build_405.dm_build_5.Dm_build_1338() == 1
	case Dm_build_362:
		dm_build_528.Schema = dm_build_526.dm_build_405.dm_build_5.Dm_build_1386(dm_build_528.getServerEncoding(), dm_build_528)
	case Dm_build_359:
		dm_build_527.explain = dm_build_526.dm_build_405.dm_build_5.Dm_build_1386(dm_build_528.getServerEncoding(), dm_build_528)

	case Dm_build_363, Dm_build_365, Dm_build_364:
		if dm_build_535 {

			counts := dm_build_526.dm_build_405.dm_build_5.Dm_build_1344()
			rowCounts := make([]int64, counts)
			for i := 0; i < int(counts); i++ {
				rowCounts[i] = dm_build_526.dm_build_405.dm_build_5.Dm_build_1347()
			}
			dm_build_527.updateCounts = rowCounts
		}

		if dm_build_536 {
			rows := dm_build_526.dm_build_405.dm_build_5.Dm_build_1344()

			var lastInsertId int64
			for i := 0; i < int(rows); i++ {
				lastInsertId = dm_build_526.dm_build_405.dm_build_5.Dm_build_1347()
			}
			dm_build_527.lastInsertId = lastInsertId

		} else if dm_build_527.updateCount == 1 {
			dm_build_527.lastInsertId = dm_build_527.rowid
		}

		if dm_build_526.dm_build_407 == EC_BP_WITH_ERROR.ErrCode {
			dm_build_526.dm_build_562(dm_build_527.updateCounts)
		}
	case Dm_build_372:
		len := dm_build_526.dm_build_405.dm_build_5.Dm_build_1356()
		dm_build_528.OracleDateFormat = dm_build_526.dm_build_405.dm_build_5.Dm_build_1381(int(len), dm_build_528.getServerEncoding(), dm_build_528)
	case Dm_build_374:

		len := dm_build_526.dm_build_405.dm_build_5.Dm_build_1356()
		dm_build_528.OracleTimestampFormat = dm_build_526.dm_build_405.dm_build_5.Dm_build_1381(int(len), dm_build_528.getServerEncoding(), dm_build_528)
	case Dm_build_375:

		len := dm_build_526.dm_build_405.dm_build_5.Dm_build_1356()
		dm_build_528.OracleTimestampTZFormat = dm_build_526.dm_build_405.dm_build_5.Dm_build_1381(int(len), dm_build_528.getServerEncoding(), dm_build_528)
	case Dm_build_373:
		len := dm_build_526.dm_build_405.dm_build_5.Dm_build_1356()
		dm_build_528.OracleTimeFormat = dm_build_526.dm_build_405.dm_build_5.Dm_build_1381(int(len), dm_build_528.getServerEncoding(), dm_build_528)
	case Dm_build_376:
		len := dm_build_526.dm_build_405.dm_build_5.Dm_build_1356()
		dm_build_528.OracleTimeTZFormat = dm_build_526.dm_build_405.dm_build_5.Dm_build_1381(int(len), dm_build_528.getServerEncoding(), dm_build_528)
	case Dm_build_377:
		dm_build_528.OracleDateLanguage = dm_build_526.dm_build_405.dm_build_5.Dm_build_1356()
	}

	return &dm_build_527, nil
}

func (dm_build_538 *dm_build_496) dm_build_537(dm_build_539 int) [][]byte {
	dm_build_540 := make([]int, dm_build_539)

	dm_build_541 := 0
	for i := 0; i < len(dm_build_538.dm_build_498); i++ {
		if dm_build_538.dm_build_498[i].ioType == IO_TYPE_INOUT || dm_build_538.dm_build_498[i].ioType == IO_TYPE_OUT {
			dm_build_540[dm_build_541] = i
			dm_build_541++
		}
	}

	dm_build_542 := make([][]byte, len(dm_build_538.dm_build_498))
	var dm_build_543 int32
	var dm_build_544 bool
	var dm_build_545 []byte = nil
	for i := 0; i < dm_build_539; i++ {
		dm_build_544 = false
		dm_build_543 = int32(dm_build_538.dm_build_405.dm_build_5.Dm_build_1359())

		if dm_build_543 == int32(Dm_build_316) {
			dm_build_543 = 0
			dm_build_544 = true
		} else if dm_build_543 == int32(Dm_build_317) {
			dm_build_543 = dm_build_538.dm_build_405.dm_build_5.Dm_build_1344()
		}

		if dm_build_544 {
			dm_build_542[dm_build_540[i]] = nil
		} else {
			dm_build_545 = dm_build_538.dm_build_405.dm_build_5.Dm_build_1365(int(dm_build_543))
			dm_build_542[dm_build_540[i]] = dm_build_545
		}
	}

	return dm_build_542
}

func (dm_build_547 *dm_build_496) dm_build_546(dm_build_548 int, dm_build_549 bool) []column {
	dm_build_550 := dm_build_547.dm_build_405.dm_build_6.getServerEncoding()
	var dm_build_551, dm_build_552, dm_build_553, dm_build_554 int16
	dm_build_555 := make([]column, dm_build_548)
	for i := 0; i < dm_build_548; i++ {
		dm_build_555[i].InitColumn()

		dm_build_555[i].colType = dm_build_547.dm_build_405.dm_build_5.Dm_build_1344()

		dm_build_555[i].prec = dm_build_547.dm_build_405.dm_build_5.Dm_build_1344()

		dm_build_555[i].scale = dm_build_547.dm_build_405.dm_build_5.Dm_build_1344()

		dm_build_555[i].nullable = dm_build_547.dm_build_405.dm_build_5.Dm_build_1344() != 0

		itemFlag := dm_build_547.dm_build_405.dm_build_5.Dm_build_1341()
		dm_build_555[i].lob = int(itemFlag)&Dm_build_493 != 0
		dm_build_555[i].identity = int(itemFlag)&Dm_build_492 != 0
		dm_build_555[i].readonly = int(itemFlag)&Dm_build_494 != 0

		dm_build_547.dm_build_405.dm_build_5.Dm_build_1244(4, false, true)

		dm_build_547.dm_build_405.dm_build_5.Dm_build_1244(2, false, true)

		dm_build_551 = dm_build_547.dm_build_405.dm_build_5.Dm_build_1341()

		dm_build_552 = dm_build_547.dm_build_405.dm_build_5.Dm_build_1341()

		dm_build_553 = dm_build_547.dm_build_405.dm_build_5.Dm_build_1341()

		dm_build_554 = dm_build_547.dm_build_405.dm_build_5.Dm_build_1341()
		dm_build_555[i].name = dm_build_547.dm_build_405.dm_build_5.Dm_build_1381(int(dm_build_551), dm_build_550, dm_build_547.dm_build_405.dm_build_6)
		dm_build_555[i].typeName = dm_build_547.dm_build_405.dm_build_5.Dm_build_1381(int(dm_build_552), dm_build_550, dm_build_547.dm_build_405.dm_build_6)
		dm_build_555[i].tableName = dm_build_547.dm_build_405.dm_build_5.Dm_build_1381(int(dm_build_553), dm_build_550, dm_build_547.dm_build_405.dm_build_6)
		dm_build_555[i].schemaName = dm_build_547.dm_build_405.dm_build_5.Dm_build_1381(int(dm_build_554), dm_build_550, dm_build_547.dm_build_405.dm_build_6)

		if dm_build_547.dm_build_408.readBaseColName {
			dm_build_555[i].baseName = dm_build_547.dm_build_405.dm_build_5.Dm_build_1394(dm_build_550, dm_build_547.dm_build_405.dm_build_6)
		}

		if dm_build_555[i].lob {
			dm_build_555[i].lobTabId = dm_build_547.dm_build_405.dm_build_5.Dm_build_1344()
			dm_build_555[i].lobColId = dm_build_547.dm_build_405.dm_build_5.Dm_build_1341()
		}

	}

	for i := 0; i < dm_build_548; i++ {

		if isComplexType(int(dm_build_555[i].colType), int(dm_build_555[i].scale)) {
			strDesc := newTypeDescriptor(dm_build_547.dm_build_405.dm_build_6)
			strDesc.unpack(dm_build_547.dm_build_405.dm_build_5)
			dm_build_555[i].typeDescriptor = strDesc
		}
	}

	return dm_build_555
}

func (dm_build_557 *dm_build_496) dm_build_556(dm_build_558 *execInfo, dm_build_559 int, dm_build_560 int, dm_build_561 int) {
	if dm_build_560 > 0 {
		startOffset := dm_build_557.dm_build_405.dm_build_5.Dm_build_1239()
		if dm_build_558.rsBdta {
			dm_build_558.rsDatas = dm_build_557.dm_build_575(dm_build_557.dm_build_408.columns, dm_build_561)
		} else {
			datas := make([][][]byte, dm_build_560)

			for i := 0; i < dm_build_560; i++ {

				datas[i] = make([][]byte, dm_build_559+1)

				dm_build_557.dm_build_405.dm_build_5.Dm_build_1244(2, false, true)

				datas[i][0] = dm_build_557.dm_build_405.dm_build_5.Dm_build_1365(LINT64_SIZE)

				dm_build_557.dm_build_405.dm_build_5.Dm_build_1244(2*dm_build_559, false, true)

				for j := 1; j < dm_build_559+1; j++ {

					colLen := dm_build_557.dm_build_405.dm_build_5.Dm_build_1359()
					if colLen == Dm_build_320 {
						datas[i][j] = nil
					} else if colLen != Dm_build_321 {
						datas[i][j] = dm_build_557.dm_build_405.dm_build_5.Dm_build_1365(int(colLen))
					} else {
						datas[i][j] = dm_build_557.dm_build_405.dm_build_5.Dm_build_1369()
					}
				}
			}

			dm_build_558.rsDatas = datas
		}
		dm_build_558.rsSizeof = dm_build_557.dm_build_405.dm_build_5.Dm_build_1239() - startOffset
	}

	if dm_build_558.rsCacheOffset > 0 {
		tbCount := dm_build_557.dm_build_405.dm_build_5.Dm_build_1341()

		ids := make([]int32, tbCount)
		tss := make([]int64, tbCount)

		for i := 0; i < int(tbCount); i++ {
			ids[i] = dm_build_557.dm_build_405.dm_build_5.Dm_build_1344()
			tss[i] = dm_build_557.dm_build_405.dm_build_5.Dm_build_1347()
		}

		dm_build_558.tbIds = ids[:]
		dm_build_558.tbTss = tss[:]
	}
}

func (dm_build_563 *dm_build_496) dm_build_562(dm_build_564 []int64) error {

	dm_build_563.dm_build_405.dm_build_5.Dm_build_1244(4, false, true)

	dm_build_565 := dm_build_563.dm_build_405.dm_build_5.Dm_build_1344()

	dm_build_566 := make([]string, 0, 8)
	for i := 0; i < int(dm_build_565); i++ {
		irow := dm_build_563.dm_build_405.dm_build_5.Dm_build_1344()

		dm_build_564[irow] = -3

		code := dm_build_563.dm_build_405.dm_build_5.Dm_build_1344()

		errStr := dm_build_563.dm_build_405.dm_build_5.Dm_build_1394(dm_build_563.dm_build_405.dm_build_6.getServerEncoding(), dm_build_563.dm_build_405.dm_build_6)

		dm_build_566 = append(dm_build_566, "row["+strconv.Itoa(int(irow))+"]:"+strconv.Itoa(int(code))+", "+errStr)
	}

	if len(dm_build_566) > 0 {
		builder := &strings.Builder{}
		for _, str := range dm_build_566 {
			builder.WriteString(util.LINE_SEPARATOR)
			builder.WriteString(str)
		}
		EC_BP_WITH_ERROR.ErrText += builder.String()
		return EC_BP_WITH_ERROR.throw()
	}
	return nil
}

const (
	Dm_build_567 = 0

	Dm_build_568 = Dm_build_567 + ULINT_SIZE

	Dm_build_569 = Dm_build_568 + USINT_SIZE

	Dm_build_570 = Dm_build_569 + ULINT_SIZE

	Dm_build_571 = Dm_build_570 + ULINT_SIZE

	Dm_build_572 = Dm_build_571 + BYTE_SIZE

	Dm_build_573 = -2

	Dm_build_574 = -3
)

func (dm_build_576 *dm_build_496) dm_build_575(dm_build_577 []column, dm_build_578 int) [][][]byte {

	dm_build_579 := dm_build_576.dm_build_405.dm_build_5.Dm_build_1362()

	dm_build_580 := dm_build_576.dm_build_405.dm_build_5.Dm_build_1359()

	var dm_build_581 bool

	if dm_build_578 >= 0 && int(dm_build_580) == len(dm_build_577)+1 {
		dm_build_581 = true
	} else {
		dm_build_581 = false
	}

	dm_build_576.dm_build_405.dm_build_5.Dm_build_1244(ULINT_SIZE, false, true)

	dm_build_576.dm_build_405.dm_build_5.Dm_build_1244(ULINT_SIZE, false, true)

	dm_build_576.dm_build_405.dm_build_5.Dm_build_1244(BYTE_SIZE, false, true)

	dm_build_582 := make([]uint16, dm_build_580)
	for icol := 0; icol < int(dm_build_580); icol++ {
		dm_build_582[icol] = dm_build_576.dm_build_405.dm_build_5.Dm_build_1359()
	}

	dm_build_583 := make([]uint32, dm_build_580)
	dm_build_584 := make([][][]byte, dm_build_579)

	for i := uint32(0); i < dm_build_579; i++ {
		dm_build_584[i] = make([][]byte, len(dm_build_577)+1)
	}

	for icol := 0; icol < int(dm_build_580); icol++ {
		dm_build_583[icol] = dm_build_576.dm_build_405.dm_build_5.Dm_build_1362()
	}

	for icol := 0; icol < int(dm_build_580); icol++ {

		dataCol := icol + 1
		if dm_build_581 && icol == dm_build_578 {
			dataCol = 0
		} else if dm_build_581 && icol > dm_build_578 {
			dataCol = icol
		}

		allNotNull := dm_build_576.dm_build_405.dm_build_5.Dm_build_1344() == 1
		var isNull []bool = nil
		if !allNotNull {
			isNull = make([]bool, dm_build_579)
			for irow := uint32(0); irow < dm_build_579; irow++ {
				isNull[irow] = dm_build_576.dm_build_405.dm_build_5.Dm_build_1338() == 0
			}
		}

		for irow := uint32(0); irow < dm_build_579; irow++ {
			if allNotNull || !isNull[irow] {
				dm_build_584[irow][dataCol] = dm_build_576.dm_build_585(int(dm_build_582[icol]))
			}
		}
	}

	if !dm_build_581 && dm_build_578 >= 0 {
		for irow := uint32(0); irow < dm_build_579; irow++ {
			dm_build_584[irow][0] = dm_build_584[irow][dm_build_578+1]
		}
	}

	return dm_build_584
}

func (dm_build_586 *dm_build_496) dm_build_585(dm_build_587 int) []byte {

	dm_build_588 := dm_build_586.dm_build_591(dm_build_587)

	dm_build_589 := int32(0)
	if dm_build_588 == Dm_build_573 {
		dm_build_589 = dm_build_586.dm_build_405.dm_build_5.Dm_build_1344()
		dm_build_588 = int(dm_build_586.dm_build_405.dm_build_5.Dm_build_1344())
	} else if dm_build_588 == Dm_build_574 {
		dm_build_588 = int(dm_build_586.dm_build_405.dm_build_5.Dm_build_1344())
	}

	dm_build_590 := dm_build_586.dm_build_405.dm_build_5.Dm_build_1365(dm_build_588 + int(dm_build_589))
	if dm_build_589 == 0 {
		return dm_build_590
	}

	for i := dm_build_588; i < len(dm_build_590); i++ {
		dm_build_590[i] = ' '
	}
	return dm_build_590
}

func (dm_build_592 *dm_build_496) dm_build_591(dm_build_593 int) int {

	dm_build_594 := 0
	switch dm_build_593 {
	case INT:
	case BIT:
	case TINYINT:
	case SMALLINT:
	case BOOLEAN:
	case NULL:
		dm_build_594 = 4

	case BIGINT:

		dm_build_594 = 8

	case CHAR:
	case VARCHAR2:
	case VARCHAR:
	case BINARY:
	case VARBINARY:
	case BLOB:
	case CLOB:
		dm_build_594 = Dm_build_573

	case DECIMAL:
		dm_build_594 = Dm_build_574

	case REAL:
		dm_build_594 = 4

	case DOUBLE:
		dm_build_594 = 8

	case DATE:
	case TIME:
	case DATETIME:
	case TIME_TZ:
	case DATETIME_TZ:
		dm_build_594 = 12

	case INTERVAL_YM:
		dm_build_594 = 12

	case INTERVAL_DT:
		dm_build_594 = 24

	default:
		panic(ECGO_INVALID_COLUMN_TYPE)
	}
	return dm_build_594
}

const (
	Dm_build_595 = Dm_build_301

	Dm_build_596 = Dm_build_595 + DDWORD_SIZE

	Dm_build_597 = Dm_build_596 + LINT64_SIZE

	Dm_build_598 = Dm_build_597 + USINT_SIZE

	Dm_build_599 = Dm_build_301

	Dm_build_600 = Dm_build_599 + DDWORD_SIZE
)

type dm_build_601 struct {
	dm_build_496
	dm_build_602 *innerRows
	dm_build_603 int64
	dm_build_604 int64
}

func dm_build_605(dm_build_606 *dm_build_2, dm_build_607 *innerRows, dm_build_608 int64, dm_build_609 int64) *dm_build_601 {
	dm_build_610 := new(dm_build_601)
	dm_build_610.dm_build_413(dm_build_606, Dm_build_279, dm_build_607.dmStmt)
	dm_build_610.dm_build_602 = dm_build_607
	dm_build_610.dm_build_603 = dm_build_608
	dm_build_610.dm_build_604 = dm_build_609
	return dm_build_610
}

func (dm_build_612 *dm_build_601) dm_build_391() error {

	dm_build_612.dm_build_405.dm_build_5.Dm_build_1414(Dm_build_595, dm_build_612.dm_build_603)

	dm_build_612.dm_build_405.dm_build_5.Dm_build_1414(Dm_build_596, dm_build_612.dm_build_604)

	dm_build_612.dm_build_405.dm_build_5.Dm_build_1406(Dm_build_597, dm_build_612.dm_build_602.id)

	dm_build_613 := dm_build_612.dm_build_602.dmStmt.dmConn.dmConnector.bufPrefetch
	var dm_build_614 int32
	if dm_build_612.dm_build_602.sizeOfRow != 0 && dm_build_612.dm_build_602.fetchSize != 0 {
		if dm_build_612.dm_build_602.sizeOfRow*dm_build_612.dm_build_602.fetchSize > int(INT32_MAX) {
			dm_build_614 = INT32_MAX
		} else {
			dm_build_614 = int32(dm_build_612.dm_build_602.sizeOfRow * dm_build_612.dm_build_602.fetchSize)
		}

		if dm_build_614 < Dm_build_332 {
			dm_build_613 = int(Dm_build_332)
		} else if dm_build_614 > Dm_build_333 {
			dm_build_613 = int(Dm_build_333)
		} else {
			dm_build_613 = int(dm_build_614)
		}

		dm_build_612.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_598, int32(dm_build_613))
	}

	return nil
}

func (dm_build_616 *dm_build_601) dm_build_395() (interface{}, error) {
	dm_build_617 := execInfo{}
	dm_build_617.rsBdta = dm_build_616.dm_build_602.isBdta

	dm_build_617.updateCount = dm_build_616.dm_build_405.dm_build_5.Dm_build_1491(Dm_build_599)
	dm_build_618 := dm_build_616.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_600)

	dm_build_616.dm_build_556(&dm_build_617, len(dm_build_616.dm_build_602.columns), int(dm_build_618), -1)

	return &dm_build_617, nil
}

type dm_build_619 struct {
	dm_build_404
	dm_build_620 *lob
	dm_build_621 int
	dm_build_622 int
}

func dm_build_623(dm_build_624 *dm_build_2, dm_build_625 *lob, dm_build_626 int, dm_build_627 int) *dm_build_619 {
	dm_build_628 := new(dm_build_619)
	dm_build_628.dm_build_409(dm_build_624, Dm_build_292)
	dm_build_628.dm_build_620 = dm_build_625
	dm_build_628.dm_build_621 = dm_build_626
	dm_build_628.dm_build_622 = dm_build_627
	return dm_build_628
}

func (dm_build_630 *dm_build_619) dm_build_391() error {

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1262(byte(dm_build_630.dm_build_620.lobFlag))

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1270(dm_build_630.dm_build_620.tabId)

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1266(dm_build_630.dm_build_620.colId)

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1286(uint64(dm_build_630.dm_build_620.blobId))

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1266(dm_build_630.dm_build_620.groupId)

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1266(dm_build_630.dm_build_620.fileId)

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1270(dm_build_630.dm_build_620.pageNo)

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1266(dm_build_630.dm_build_620.curFileId)

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1270(dm_build_630.dm_build_620.curPageNo)

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1270(dm_build_630.dm_build_620.totalOffset)

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1270(int32(dm_build_630.dm_build_621))

	dm_build_630.dm_build_405.dm_build_5.Dm_build_1270(int32(dm_build_630.dm_build_622))

	if dm_build_630.dm_build_405.dm_build_6.NewLobFlag {
		dm_build_630.dm_build_405.dm_build_5.Dm_build_1286(uint64(dm_build_630.dm_build_620.rowId))
		dm_build_630.dm_build_405.dm_build_5.Dm_build_1266(dm_build_630.dm_build_620.exGroupId)
		dm_build_630.dm_build_405.dm_build_5.Dm_build_1266(dm_build_630.dm_build_620.exFileId)
		dm_build_630.dm_build_405.dm_build_5.Dm_build_1270(dm_build_630.dm_build_620.exPageNo)
	}

	return nil
}

func (dm_build_632 *dm_build_619) dm_build_395() (interface{}, error) {

	dm_build_632.dm_build_620.readOver = dm_build_632.dm_build_405.dm_build_5.Dm_build_1338() == 1
	var dm_build_633 = dm_build_632.dm_build_405.dm_build_5.Dm_build_1362()
	if dm_build_633 <= 0 {
		return []byte(nil), nil
	}
	dm_build_632.dm_build_620.curFileId = dm_build_632.dm_build_405.dm_build_5.Dm_build_1341()
	dm_build_632.dm_build_620.curPageNo = dm_build_632.dm_build_405.dm_build_5.Dm_build_1344()
	dm_build_632.dm_build_620.totalOffset = dm_build_632.dm_build_405.dm_build_5.Dm_build_1344()

	return dm_build_632.dm_build_405.dm_build_5.Dm_build_1375(int(dm_build_633)), nil
}

type dm_build_634 struct {
	dm_build_404
	dm_build_635 *lob
}

func dm_build_636(dm_build_637 *dm_build_2, dm_build_638 *lob) *dm_build_634 {
	dm_build_639 := new(dm_build_634)
	dm_build_639.dm_build_409(dm_build_637, Dm_build_289)
	dm_build_639.dm_build_635 = dm_build_638
	return dm_build_639
}

func (dm_build_641 *dm_build_634) dm_build_391() error {

	dm_build_641.dm_build_405.dm_build_5.Dm_build_1262(byte(dm_build_641.dm_build_635.lobFlag))

	dm_build_641.dm_build_405.dm_build_5.Dm_build_1286(uint64(dm_build_641.dm_build_635.blobId))

	dm_build_641.dm_build_405.dm_build_5.Dm_build_1266(dm_build_641.dm_build_635.groupId)

	dm_build_641.dm_build_405.dm_build_5.Dm_build_1266(dm_build_641.dm_build_635.fileId)

	dm_build_641.dm_build_405.dm_build_5.Dm_build_1270(dm_build_641.dm_build_635.pageNo)

	if dm_build_641.dm_build_405.dm_build_6.NewLobFlag {
		dm_build_641.dm_build_405.dm_build_5.Dm_build_1270(dm_build_641.dm_build_635.tabId)
		dm_build_641.dm_build_405.dm_build_5.Dm_build_1266(dm_build_641.dm_build_635.colId)
		dm_build_641.dm_build_405.dm_build_5.Dm_build_1286(uint64(dm_build_641.dm_build_635.rowId))

		dm_build_641.dm_build_405.dm_build_5.Dm_build_1266(dm_build_641.dm_build_635.exGroupId)
		dm_build_641.dm_build_405.dm_build_5.Dm_build_1266(dm_build_641.dm_build_635.exFileId)
		dm_build_641.dm_build_405.dm_build_5.Dm_build_1270(dm_build_641.dm_build_635.exPageNo)
	}

	return nil
}

func (dm_build_643 *dm_build_634) dm_build_395() (interface{}, error) {

	return int64(dm_build_643.dm_build_405.dm_build_5.Dm_build_1344()), nil
}

type dm_build_644 struct {
	dm_build_404
	dm_build_645 *lob
	dm_build_646 int
}

func dm_build_647(dm_build_648 *dm_build_2, dm_build_649 *lob, dm_build_650 int) *dm_build_644 {
	dm_build_651 := new(dm_build_644)
	dm_build_651.dm_build_409(dm_build_648, Dm_build_291)
	dm_build_651.dm_build_645 = dm_build_649
	dm_build_651.dm_build_646 = dm_build_650
	return dm_build_651
}

func (dm_build_653 *dm_build_644) dm_build_391() error {

	dm_build_653.dm_build_405.dm_build_5.Dm_build_1262(byte(dm_build_653.dm_build_645.lobFlag))

	dm_build_653.dm_build_405.dm_build_5.Dm_build_1286(uint64(dm_build_653.dm_build_645.blobId))

	dm_build_653.dm_build_405.dm_build_5.Dm_build_1266(dm_build_653.dm_build_645.groupId)

	dm_build_653.dm_build_405.dm_build_5.Dm_build_1266(dm_build_653.dm_build_645.fileId)

	dm_build_653.dm_build_405.dm_build_5.Dm_build_1270(dm_build_653.dm_build_645.pageNo)

	dm_build_653.dm_build_405.dm_build_5.Dm_build_1270(dm_build_653.dm_build_645.tabId)
	dm_build_653.dm_build_405.dm_build_5.Dm_build_1266(dm_build_653.dm_build_645.colId)
	dm_build_653.dm_build_405.dm_build_5.Dm_build_1286(uint64(dm_build_653.dm_build_645.rowId))
	dm_build_653.dm_build_405.dm_build_5.Dm_build_1298(Dm_build_865.Dm_build_1064(uint32(dm_build_653.dm_build_646)))

	if dm_build_653.dm_build_405.dm_build_6.NewLobFlag {
		dm_build_653.dm_build_405.dm_build_5.Dm_build_1266(dm_build_653.dm_build_645.exGroupId)
		dm_build_653.dm_build_405.dm_build_5.Dm_build_1266(dm_build_653.dm_build_645.exFileId)
		dm_build_653.dm_build_405.dm_build_5.Dm_build_1270(dm_build_653.dm_build_645.exPageNo)
	}
	return nil
}

func (dm_build_655 *dm_build_644) dm_build_395() (interface{}, error) {

	dm_build_656 := dm_build_655.dm_build_405.dm_build_5.Dm_build_1362()
	dm_build_655.dm_build_645.blobId = dm_build_655.dm_build_405.dm_build_5.Dm_build_1347()
	dm_build_655.dm_build_645.resetCurrentInfo()
	return int64(dm_build_656), nil
}

const (
	Dm_build_657 = Dm_build_301

	Dm_build_658 = Dm_build_657 + ULINT_SIZE

	Dm_build_659 = Dm_build_658 + ULINT_SIZE

	Dm_build_660 = Dm_build_659 + ULINT_SIZE

	Dm_build_661 = Dm_build_660 + BYTE_SIZE

	Dm_build_662 = Dm_build_661 + USINT_SIZE

	Dm_build_663 = Dm_build_662 + ULINT_SIZE

	Dm_build_664 = Dm_build_663 + BYTE_SIZE

	Dm_build_665 = Dm_build_664 + BYTE_SIZE

	Dm_build_666 = Dm_build_665 + BYTE_SIZE

	Dm_build_667 = Dm_build_301

	Dm_build_668 = Dm_build_667 + ULINT_SIZE

	Dm_build_669 = Dm_build_668 + ULINT_SIZE

	Dm_build_670 = Dm_build_669 + BYTE_SIZE

	Dm_build_671 = Dm_build_670 + ULINT_SIZE

	Dm_build_672 = Dm_build_671 + BYTE_SIZE

	Dm_build_673 = Dm_build_672 + BYTE_SIZE

	Dm_build_674 = Dm_build_673 + USINT_SIZE

	Dm_build_675 = Dm_build_674 + USINT_SIZE

	Dm_build_676 = Dm_build_675 + BYTE_SIZE

	Dm_build_677 = Dm_build_676 + USINT_SIZE

	Dm_build_678 = Dm_build_677 + BYTE_SIZE

	Dm_build_679 = Dm_build_678 + BYTE_SIZE

	Dm_build_680 = Dm_build_679 + ULINT_SIZE
)

type dm_build_681 struct {
	dm_build_404

	dm_build_682 *DmConnection

	dm_build_683 bool
}

func dm_build_684(dm_build_685 *dm_build_2) *dm_build_681 {
	dm_build_686 := new(dm_build_681)
	dm_build_686.dm_build_409(dm_build_685, Dm_build_273)
	dm_build_686.dm_build_682 = dm_build_685.dm_build_6
	return dm_build_686
}

func (dm_build_688 *dm_build_681) dm_build_391() error {

	dm_build_688.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_657, Dm_build_312)

	dm_build_688.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_658, g2dbIsoLevel(dm_build_688.dm_build_682.IsoLevel))
	dm_build_688.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_659, int32(Locale))
	dm_build_688.dm_build_405.dm_build_5.Dm_build_1406(Dm_build_661, dm_build_688.dm_build_682.dmConnector.localTimezone)

	if dm_build_688.dm_build_682.ReadOnly {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_660, Dm_build_335)
	} else {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_660, Dm_build_334)
	}

	dm_build_688.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_662, int32(dm_build_688.dm_build_682.dmConnector.sessionTimeout))

	if dm_build_688.dm_build_682.dmConnector.mppLocal {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_663, 1)
	} else {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_663, 0)
	}

	if dm_build_688.dm_build_682.dmConnector.rwSeparate {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_664, 1)
	} else {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_664, 0)
	}

	if dm_build_688.dm_build_682.NewLobFlag {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_665, 1)
	} else {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_665, 0)
	}

	dm_build_688.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_666, dm_build_688.dm_build_682.dmConnector.osAuthType)

	dm_build_689 := dm_build_688.dm_build_682.getServerEncoding()

	if dm_build_688.dm_build_405.dm_build_11 != "" {

	}

	dm_build_690 := Dm_build_865.Dm_build_1075(dm_build_688.dm_build_682.dmConnector.user, dm_build_689, dm_build_688.dm_build_405.dm_build_6)
	dm_build_691 := Dm_build_865.Dm_build_1075(dm_build_688.dm_build_682.dmConnector.password, dm_build_689, dm_build_688.dm_build_405.dm_build_6)
	if len(dm_build_690) > Dm_build_309 {
		return ECGO_USERNAME_TOO_LONG.throw()
	}
	if len(dm_build_691) > Dm_build_309 {
		return ECGO_PASSWORD_TOO_LONG.throw()
	}

	if dm_build_688.dm_build_405.dm_build_8 && dm_build_688.dm_build_682.dmConnector.loginCertificate != "" {

	} else if dm_build_688.dm_build_405.dm_build_8 {
		dm_build_690 = dm_build_688.dm_build_405.dm_build_7.Encrypt(dm_build_690, false)
		dm_build_691 = dm_build_688.dm_build_405.dm_build_7.Encrypt(dm_build_691, false)
	}

	dm_build_688.dm_build_405.dm_build_5.Dm_build_1302(dm_build_690)
	dm_build_688.dm_build_405.dm_build_5.Dm_build_1302(dm_build_691)

	dm_build_688.dm_build_405.dm_build_5.Dm_build_1314(dm_build_688.dm_build_682.dmConnector.appName, dm_build_689, dm_build_688.dm_build_405.dm_build_6)
	dm_build_688.dm_build_405.dm_build_5.Dm_build_1314(dm_build_688.dm_build_682.dmConnector.osName, dm_build_689, dm_build_688.dm_build_405.dm_build_6)

	if hostName, err := os.Hostname(); err != nil {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1314(hostName, dm_build_689, dm_build_688.dm_build_405.dm_build_6)
	} else {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1314("", dm_build_689, dm_build_688.dm_build_405.dm_build_6)
	}

	if dm_build_688.dm_build_682.dmConnector.rwStandby {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1262(1)
	} else {
		dm_build_688.dm_build_405.dm_build_5.Dm_build_1262(0)
	}

	return nil
}

func (dm_build_693 *dm_build_681) dm_build_395() (interface{}, error) {

	dm_build_693.dm_build_682.MaxRowSize = dm_build_693.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_667)
	dm_build_693.dm_build_682.DDLAutoCommit = dm_build_693.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_669) == 1
	dm_build_693.dm_build_682.IsoLevel = dm_build_693.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_670)
	dm_build_693.dm_build_682.dmConnector.caseSensitive = dm_build_693.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_671) == 1
	dm_build_693.dm_build_682.BackslashEscape = dm_build_693.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_672) == 1
	dm_build_693.dm_build_682.SvrStat = dm_build_693.dm_build_405.dm_build_5.Dm_build_1485(Dm_build_674)
	dm_build_693.dm_build_682.SvrMode = dm_build_693.dm_build_405.dm_build_5.Dm_build_1485(Dm_build_673)
	dm_build_693.dm_build_682.ConstParaOpt = dm_build_693.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_675) == 1
	dm_build_693.dm_build_682.DbTimezone = dm_build_693.dm_build_405.dm_build_5.Dm_build_1485(Dm_build_676)
	dm_build_693.dm_build_682.NewLobFlag = dm_build_693.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_678) == 1

	if dm_build_693.dm_build_682.dmConnector.bufPrefetch == 0 {
		dm_build_693.dm_build_682.dmConnector.bufPrefetch = int(dm_build_693.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_679))
	}

	dm_build_693.dm_build_682.LifeTimeRemainder = dm_build_693.dm_build_405.dm_build_5.Dm_build_1485(Dm_build_680)

	dm_build_694 := dm_build_693.dm_build_682.getServerEncoding()

	dm_build_693.dm_build_682.InstanceName = dm_build_693.dm_build_405.dm_build_5.Dm_build_1386(dm_build_694, dm_build_693.dm_build_405.dm_build_6)
	dm_build_693.dm_build_682.Schema = dm_build_693.dm_build_405.dm_build_5.Dm_build_1386(dm_build_694, dm_build_693.dm_build_405.dm_build_6)
	dm_build_693.dm_build_682.LastLoginIP = dm_build_693.dm_build_405.dm_build_5.Dm_build_1386(dm_build_694, dm_build_693.dm_build_405.dm_build_6)
	dm_build_693.dm_build_682.LastLoginTime = dm_build_693.dm_build_405.dm_build_5.Dm_build_1386(dm_build_694, dm_build_693.dm_build_405.dm_build_6)
	dm_build_693.dm_build_682.FailedAttempts = dm_build_693.dm_build_405.dm_build_5.Dm_build_1344()
	dm_build_693.dm_build_682.LoginWarningID = dm_build_693.dm_build_405.dm_build_5.Dm_build_1344()
	dm_build_693.dm_build_682.GraceTimeRemainder = dm_build_693.dm_build_405.dm_build_5.Dm_build_1344()
	dm_build_693.dm_build_682.Guid = dm_build_693.dm_build_405.dm_build_5.Dm_build_1386(dm_build_694, dm_build_693.dm_build_405.dm_build_6)
	dm_build_693.dm_build_682.DbName = dm_build_693.dm_build_405.dm_build_5.Dm_build_1386(dm_build_694, dm_build_693.dm_build_405.dm_build_6)

	if dm_build_693.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_677) == 1 {
		dm_build_693.dm_build_682.StandbyHost = dm_build_693.dm_build_405.dm_build_5.Dm_build_1386(dm_build_694, dm_build_693.dm_build_405.dm_build_6)
		dm_build_693.dm_build_682.StandbyPort = dm_build_693.dm_build_405.dm_build_5.Dm_build_1344()
		dm_build_693.dm_build_682.StandbyCount = dm_build_693.dm_build_405.dm_build_5.Dm_build_1359()
	}

	if dm_build_693.dm_build_405.dm_build_5.Dm_build_1241(false) > 0 {
		dm_build_693.dm_build_682.SessionID = dm_build_693.dm_build_405.dm_build_5.Dm_build_1347()
	}

	if dm_build_693.dm_build_405.dm_build_5.Dm_build_1241(false) > 0 {
		if dm_build_693.dm_build_405.dm_build_5.Dm_build_1338() == 1 {

			dm_build_693.dm_build_682.OracleDateFormat = "DD-MON-YY"

			dm_build_693.dm_build_682.OracleTimestampFormat = "DD-MON-YY HH12.MI.SS.FF6 AM"

			dm_build_693.dm_build_682.OracleTimestampTZFormat = "DD-MON-YY HH12.MI.SS.FF6 AM +TZH:TZM"

			dm_build_693.dm_build_682.OracleTimeFormat = "HH12.MI.SS.FF6 AM"

			dm_build_693.dm_build_682.OracleTimeTZFormat = "HH12.MI.SS.FF6 AM +TZH:TZM"
		}
	}

	return nil, nil
}

const (
	Dm_build_695 = Dm_build_301
)

type dm_build_696 struct {
	dm_build_496
	dm_build_697 int16
}

func dm_build_698(dm_build_699 *dm_build_2, dm_build_700 *DmStatement, dm_build_701 int16) *dm_build_696 {
	dm_build_702 := new(dm_build_696)
	dm_build_702.dm_build_413(dm_build_699, Dm_build_293, dm_build_700)
	dm_build_702.dm_build_697 = dm_build_701
	return dm_build_702
}

func (dm_build_704 *dm_build_696) dm_build_391() error {
	dm_build_704.dm_build_405.dm_build_5.Dm_build_1406(Dm_build_695, dm_build_704.dm_build_697)
	return nil
}

func (dm_build_706 *dm_build_696) dm_build_395() (interface{}, error) {
	return dm_build_706.dm_build_496.dm_build_395()
}

const (
	Dm_build_707 = Dm_build_301
)

type dm_build_708 struct {
	dm_build_496
	dm_build_709 []parameter
}

func dm_build_710(dm_build_711 *dm_build_2, dm_build_712 *DmStatement, dm_build_713 []parameter) *dm_build_708 {
	dm_build_714 := new(dm_build_708)
	dm_build_714.dm_build_413(dm_build_711, Dm_build_297, dm_build_712)
	dm_build_714.dm_build_709 = dm_build_713
	return dm_build_714
}

func (dm_build_716 *dm_build_708) dm_build_391() error {

	if dm_build_716.dm_build_709 == nil {
		dm_build_716.dm_build_405.dm_build_5.Dm_build_1406(Dm_build_707, 0)
	} else {
		dm_build_716.dm_build_405.dm_build_5.Dm_build_1406(Dm_build_707, int16(len(dm_build_716.dm_build_709)))
	}

	return dm_build_716.dm_build_519(dm_build_716.dm_build_709)
}

type dm_build_717 struct {
	dm_build_496
	dm_build_718 bool
	dm_build_719 int16
}

func dm_build_720(dm_build_721 *dm_build_2, dm_build_722 *DmStatement, dm_build_723 bool, dm_build_724 int16) *dm_build_717 {
	dm_build_725 := new(dm_build_717)
	dm_build_725.dm_build_413(dm_build_721, Dm_build_277, dm_build_722)
	dm_build_725.dm_build_718 = dm_build_723
	dm_build_725.dm_build_719 = dm_build_724
	return dm_build_725
}

func (dm_build_727 *dm_build_717) dm_build_391() error {

	dm_build_728 := Dm_build_301

	if dm_build_727.dm_build_405.dm_build_6.autoCommit {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 1)
	} else {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 0)
	}

	if dm_build_727.dm_build_718 {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 1)
	} else {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 0)
	}

	dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 0)

	dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 1)

	if dm_build_727.dm_build_405.dm_build_6.CompatibleOracle() {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 0)
	} else {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 1)
	}

	dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1406(dm_build_728, dm_build_727.dm_build_719)

	if dm_build_727.dm_build_408.maxRows <= 0 || dm_build_727.dm_build_405.dm_build_6.dmConnector.enRsCache {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1414(dm_build_728, INT64_MAX)
	} else {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1414(dm_build_728, dm_build_727.dm_build_408.maxRows)
	}

	if dm_build_727.dm_build_405.dm_build_6.dmConnector.isBdtaRS {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, Dm_build_379)
	} else {
		dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, Dm_build_378)
	}

	dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1406(dm_build_728, 0)

	dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 1)

	dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 0)

	dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1402(dm_build_728, 0)

	dm_build_728 += dm_build_727.dm_build_405.dm_build_5.Dm_build_1410(dm_build_728, dm_build_727.dm_build_408.queryTimeout)

	dm_build_727.dm_build_405.dm_build_5.Dm_build_1332(dm_build_727.dm_build_408.nativeSql, dm_build_727.dm_build_405.dm_build_6.getServerEncoding(), dm_build_727.dm_build_405.dm_build_6)

	return nil
}

func (dm_build_730 *dm_build_717) dm_build_395() (interface{}, error) {

	if dm_build_730.dm_build_718 {
		return dm_build_730.dm_build_496.dm_build_395()
	}

	dm_build_731 := NewExceInfo()
	dm_build_732 := Dm_build_301

	dm_build_731.retSqlType = dm_build_730.dm_build_405.dm_build_5.Dm_build_1485(dm_build_732)
	dm_build_732 += USINT_SIZE

	dm_build_733 := dm_build_730.dm_build_405.dm_build_5.Dm_build_1485(dm_build_732)
	dm_build_732 += USINT_SIZE

	dm_build_734 := dm_build_730.dm_build_405.dm_build_5.Dm_build_1485(dm_build_732)
	dm_build_732 += USINT_SIZE

	dm_build_730.dm_build_405.dm_build_5.Dm_build_1491(dm_build_732)
	dm_build_732 += DDWORD_SIZE

	dm_build_730.dm_build_405.dm_build_6.TrxStatus = dm_build_730.dm_build_405.dm_build_5.Dm_build_1488(dm_build_732)
	dm_build_732 += ULINT_SIZE

	if dm_build_733 > 0 {
		dm_build_730.dm_build_408.params = dm_build_730.dm_build_735(int(dm_build_733))
		dm_build_730.dm_build_408.paramCount = dm_build_733
	} else {
		dm_build_730.dm_build_408.params = make([]parameter, 0)
		dm_build_730.dm_build_408.paramCount = 0
	}

	if dm_build_734 > 0 {
		dm_build_730.dm_build_408.columns = dm_build_730.dm_build_546(int(dm_build_734), dm_build_731.rsBdta)
	} else {
		dm_build_730.dm_build_408.columns = make([]column, 0)
	}

	return dm_build_731, nil
}

func (dm_build_736 *dm_build_717) dm_build_735(dm_build_737 int) []parameter {

	var dm_build_738, dm_build_739, dm_build_740, dm_build_741 int16

	dm_build_742 := make([]parameter, dm_build_737)
	for i := 0; i < dm_build_737; i++ {

		dm_build_742[i].InitParameter()

		dm_build_742[i].colType = dm_build_736.dm_build_405.dm_build_5.Dm_build_1344()

		dm_build_742[i].prec = dm_build_736.dm_build_405.dm_build_5.Dm_build_1344()

		dm_build_742[i].scale = dm_build_736.dm_build_405.dm_build_5.Dm_build_1344()

		dm_build_742[i].nullable = dm_build_736.dm_build_405.dm_build_5.Dm_build_1344() != 0

		itemFlag := dm_build_736.dm_build_405.dm_build_5.Dm_build_1341()

		if int(itemFlag)&Dm_build_495 != 0 {
			dm_build_742[i].typeFlag = TYPE_FLAG_RECOMMEND
		} else {
			dm_build_742[i].typeFlag = TYPE_FLAG_EXACT
		}

		dm_build_742[i].lob = int(itemFlag)&Dm_build_493 != 0

		dm_build_736.dm_build_405.dm_build_5.Dm_build_1344()

		dm_build_742[i].ioType = byte(dm_build_736.dm_build_405.dm_build_5.Dm_build_1341())

		dm_build_738 = dm_build_736.dm_build_405.dm_build_5.Dm_build_1341()

		dm_build_739 = dm_build_736.dm_build_405.dm_build_5.Dm_build_1341()

		dm_build_740 = dm_build_736.dm_build_405.dm_build_5.Dm_build_1341()

		dm_build_741 = dm_build_736.dm_build_405.dm_build_5.Dm_build_1341()
		dm_build_742[i].name = dm_build_736.dm_build_405.dm_build_5.Dm_build_1381(int(dm_build_738), dm_build_736.dm_build_405.dm_build_6.getServerEncoding(), dm_build_736.dm_build_405.dm_build_6)
		dm_build_742[i].typeName = dm_build_736.dm_build_405.dm_build_5.Dm_build_1381(int(dm_build_739), dm_build_736.dm_build_405.dm_build_6.getServerEncoding(), dm_build_736.dm_build_405.dm_build_6)
		dm_build_742[i].tableName = dm_build_736.dm_build_405.dm_build_5.Dm_build_1381(int(dm_build_740), dm_build_736.dm_build_405.dm_build_6.getServerEncoding(), dm_build_736.dm_build_405.dm_build_6)
		dm_build_742[i].schemaName = dm_build_736.dm_build_405.dm_build_5.Dm_build_1381(int(dm_build_741), dm_build_736.dm_build_405.dm_build_6.getServerEncoding(), dm_build_736.dm_build_405.dm_build_6)

		if dm_build_742[i].lob {
			dm_build_742[i].lobTabId = dm_build_736.dm_build_405.dm_build_5.Dm_build_1344()
			dm_build_742[i].lobColId = dm_build_736.dm_build_405.dm_build_5.Dm_build_1341()
		}
	}

	for i := 0; i < dm_build_737; i++ {

		if isComplexType(int(dm_build_742[i].colType), int(dm_build_742[i].scale)) {

			strDesc := newTypeDescriptor(dm_build_736.dm_build_405.dm_build_6)
			strDesc.unpack(dm_build_736.dm_build_405.dm_build_5)
			dm_build_742[i].typeDescriptor = strDesc
		}
	}

	return dm_build_742
}

const (
	Dm_build_743 = Dm_build_301
)

type dm_build_744 struct {
	dm_build_404
	dm_build_745 int16
	dm_build_746 *Dm_build_1141
	dm_build_747 int32
}

func dm_build_748(dm_build_749 *dm_build_2, dm_build_750 *DmStatement, dm_build_751 int16, dm_build_752 *Dm_build_1141, dm_build_753 int32) *dm_build_744 {
	dm_build_754 := new(dm_build_744)
	dm_build_754.dm_build_413(dm_build_749, Dm_build_283, dm_build_750)
	dm_build_754.dm_build_745 = dm_build_751
	dm_build_754.dm_build_746 = dm_build_752
	dm_build_754.dm_build_747 = dm_build_753
	return dm_build_754
}

func (dm_build_756 *dm_build_744) dm_build_391() error {
	dm_build_756.dm_build_405.dm_build_5.Dm_build_1406(Dm_build_743, dm_build_756.dm_build_745)

	dm_build_756.dm_build_405.dm_build_5.Dm_build_1270(dm_build_756.dm_build_747)

	if dm_build_756.dm_build_405.dm_build_6.NewLobFlag {
		dm_build_756.dm_build_405.dm_build_5.Dm_build_1270(-1)
	}
	dm_build_756.dm_build_746.Dm_build_1148(dm_build_756.dm_build_405.dm_build_5, int(dm_build_756.dm_build_747))
	return nil
}

type dm_build_757 struct {
	dm_build_404
}

func dm_build_758(dm_build_759 *dm_build_2) *dm_build_757 {
	dm_build_760 := new(dm_build_757)
	dm_build_760.dm_build_409(dm_build_759, Dm_build_281)
	return dm_build_760
}

type dm_build_761 struct {
	dm_build_404
	dm_build_762 int32
}

func dm_build_763(dm_build_764 *dm_build_2, dm_build_765 int32) *dm_build_761 {
	dm_build_766 := new(dm_build_761)
	dm_build_766.dm_build_409(dm_build_764, Dm_build_294)
	dm_build_766.dm_build_762 = dm_build_765
	return dm_build_766
}

func (dm_build_768 *dm_build_761) dm_build_391() error {

	dm_build_769 := Dm_build_301
	dm_build_769 += dm_build_768.dm_build_405.dm_build_5.Dm_build_1410(dm_build_769, g2dbIsoLevel(dm_build_768.dm_build_762))
	return nil
}

type dm_build_770 struct {
	dm_build_404
	dm_build_771 *lob
	dm_build_772 byte
	dm_build_773 int
	dm_build_774 []byte
	dm_build_775 int
	dm_build_776 int
}

func dm_build_777(dm_build_778 *dm_build_2, dm_build_779 *lob, dm_build_780 byte, dm_build_781 int, dm_build_782 []byte,
	dm_build_783 int, dm_build_784 int) *dm_build_770 {
	dm_build_785 := new(dm_build_770)
	dm_build_785.dm_build_409(dm_build_778, Dm_build_290)
	dm_build_785.dm_build_771 = dm_build_779
	dm_build_785.dm_build_772 = dm_build_780
	dm_build_785.dm_build_773 = dm_build_781
	dm_build_785.dm_build_774 = dm_build_782
	dm_build_785.dm_build_775 = dm_build_783
	dm_build_785.dm_build_776 = dm_build_784
	return dm_build_785
}

func (dm_build_787 *dm_build_770) dm_build_391() error {

	dm_build_787.dm_build_405.dm_build_5.Dm_build_1262(byte(dm_build_787.dm_build_771.lobFlag))
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1262(dm_build_787.dm_build_772)
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1286(uint64(dm_build_787.dm_build_771.blobId))
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1266(dm_build_787.dm_build_771.groupId)
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1266(dm_build_787.dm_build_771.fileId)
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1270(dm_build_787.dm_build_771.pageNo)
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1266(dm_build_787.dm_build_771.curFileId)
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1270(dm_build_787.dm_build_771.curPageNo)
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1270(dm_build_787.dm_build_771.totalOffset)
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1270(dm_build_787.dm_build_771.tabId)
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1266(dm_build_787.dm_build_771.colId)
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1286(uint64(dm_build_787.dm_build_771.rowId))

	dm_build_787.dm_build_405.dm_build_5.Dm_build_1270(int32(dm_build_787.dm_build_773))
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1270(int32(dm_build_787.dm_build_776))
	dm_build_787.dm_build_405.dm_build_5.Dm_build_1298(dm_build_787.dm_build_774)

	if dm_build_787.dm_build_405.dm_build_6.NewLobFlag {
		dm_build_787.dm_build_405.dm_build_5.Dm_build_1266(dm_build_787.dm_build_771.exGroupId)
		dm_build_787.dm_build_405.dm_build_5.Dm_build_1266(dm_build_787.dm_build_771.exFileId)
		dm_build_787.dm_build_405.dm_build_5.Dm_build_1270(dm_build_787.dm_build_771.exPageNo)
	}
	return nil
}

func (dm_build_789 *dm_build_770) dm_build_395() (interface{}, error) {

	var dm_build_790 = dm_build_789.dm_build_405.dm_build_5.Dm_build_1344()
	dm_build_789.dm_build_771.blobId = dm_build_789.dm_build_405.dm_build_5.Dm_build_1347()
	dm_build_789.dm_build_771.fileId = dm_build_789.dm_build_405.dm_build_5.Dm_build_1341()
	dm_build_789.dm_build_771.pageNo = dm_build_789.dm_build_405.dm_build_5.Dm_build_1344()
	dm_build_789.dm_build_771.curFileId = dm_build_789.dm_build_405.dm_build_5.Dm_build_1341()
	dm_build_789.dm_build_771.curPageNo = dm_build_789.dm_build_405.dm_build_5.Dm_build_1344()
	dm_build_789.dm_build_771.totalOffset = dm_build_789.dm_build_405.dm_build_5.Dm_build_1344()
	return dm_build_790, nil
}

const (
	Dm_build_791 = Dm_build_301

	Dm_build_792 = Dm_build_791 + ULINT_SIZE

	Dm_build_793 = Dm_build_792 + ULINT_SIZE

	Dm_build_794 = Dm_build_793 + BYTE_SIZE

	Dm_build_795 = Dm_build_794 + BYTE_SIZE

	Dm_build_796 = Dm_build_795 + BYTE_SIZE

	Dm_build_797 = Dm_build_796 + BYTE_SIZE

	Dm_build_798 = Dm_build_301

	Dm_build_799 = Dm_build_798 + ULINT_SIZE

	Dm_build_800 = Dm_build_799 + ULINT_SIZE

	Dm_build_801 = Dm_build_800 + ULINT_SIZE

	Dm_build_802 = Dm_build_801 + ULINT_SIZE

	Dm_build_803 = Dm_build_802 + ULINT_SIZE

	Dm_build_804 = Dm_build_803 + BYTE_SIZE

	Dm_build_805 = Dm_build_804 + BYTE_SIZE

	Dm_build_806 = Dm_build_805 + BYTE_SIZE
)

type dm_build_807 struct {
	dm_build_404
	dm_build_808 *DmConnection
	dm_build_809 int
	Dm_build_810 int32
	Dm_build_811 []byte
	dm_build_812 byte
}

func dm_build_813(dm_build_814 *dm_build_2) *dm_build_807 {
	dm_build_815 := new(dm_build_807)
	dm_build_815.dm_build_409(dm_build_814, Dm_build_299)
	dm_build_815.dm_build_808 = dm_build_814.dm_build_6
	return dm_build_815
}

func dm_build_816(dm_build_817 string, dm_build_818 string) int {
	dm_build_819 := strings.Split(dm_build_817, ".")
	dm_build_820 := strings.Split(dm_build_818, ".")

	for i, serStr := range dm_build_819 {
		ser, _ := strconv.ParseInt(serStr, 10, 32)
		global, _ := strconv.ParseInt(dm_build_820[i], 10, 32)
		if ser < global {
			return -1
		} else if ser == global {
			continue
		} else {
			return 1
		}
	}

	return 0
}

func (dm_build_822 *dm_build_807) dm_build_391() error {

	dm_build_822.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_791, int32(0))
	dm_build_822.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_792, int32(dm_build_822.dm_build_808.dmConnector.compress))

	if dm_build_822.dm_build_808.dmConnector.loginEncrypt {
		dm_build_822.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_794, 2)
		dm_build_822.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_793, 1)
	} else {
		dm_build_822.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_794, 0)
		dm_build_822.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_793, 0)
	}

	if dm_build_822.dm_build_808.dmConnector.isBdtaRS {
		dm_build_822.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_795, Dm_build_379)
	} else {
		dm_build_822.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_795, Dm_build_378)
	}

	dm_build_822.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_796, byte(dm_build_822.dm_build_808.dmConnector.compressID))

	if dm_build_822.dm_build_808.dmConnector.loginCertificate != "" {
		dm_build_822.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_797, 1)
	} else {
		dm_build_822.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_797, 0)
	}

	dm_build_823 := dm_build_822.dm_build_808.getServerEncoding()
	dm_build_822.dm_build_405.dm_build_5.Dm_build_1314(Dm_build_266, dm_build_823, dm_build_822.dm_build_405.dm_build_6)

	var dm_build_824 byte
	if dm_build_822.dm_build_808.dmConnector.uKeyName != "" {
		dm_build_824 = 1
	} else {
		dm_build_824 = 0
	}

	dm_build_822.dm_build_405.dm_build_5.Dm_build_1262(0)

	if dm_build_824 == 1 {

	}

	if dm_build_822.dm_build_808.dmConnector.loginEncrypt {
		clientPubKey, err := dm_build_822.dm_build_405.dm_build_246()
		if err != nil {
			return err
		}
		dm_build_822.dm_build_405.dm_build_5.Dm_build_1302(clientPubKey)
	}

	return nil
}

func (dm_build_826 *dm_build_807) dm_build_394() error {
	dm_build_826.dm_build_405.dm_build_5.Dm_build_1236(0)
	dm_build_826.dm_build_405.dm_build_5.Dm_build_1244(Dm_build_300, false, true)
	return nil
}

func (dm_build_828 *dm_build_807) dm_build_395() (interface{}, error) {

	dm_build_828.dm_build_808.sslEncrypt = int(dm_build_828.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_798))
	dm_build_828.dm_build_808.GlobalServerSeries = int(dm_build_828.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_799))

	switch dm_build_828.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_800) {
	case 1:
		dm_build_828.dm_build_808.serverEncoding = ENCODING_UTF8
	case 2:
		dm_build_828.dm_build_808.serverEncoding = ENCODING_EUCKR
	default:
		dm_build_828.dm_build_808.serverEncoding = ENCODING_GB18030
	}

	dm_build_828.dm_build_808.dmConnector.compress = int(dm_build_828.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_801))
	dm_build_829 := dm_build_828.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_803)
	dm_build_830 := dm_build_828.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_804)
	dm_build_828.dm_build_808.dmConnector.isBdtaRS = dm_build_828.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_805) > 0
	dm_build_828.dm_build_808.dmConnector.compressID = int8(dm_build_828.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_806))

	dm_build_831 := dm_build_828.dm_build_437()
	if dm_build_831 != nil {
		return nil, dm_build_831
	}

	dm_build_832 := dm_build_828.dm_build_405.dm_build_5.Dm_build_1386(dm_build_828.dm_build_808.getServerEncoding(), dm_build_828.dm_build_405.dm_build_6)
	if dm_build_816(dm_build_832, Dm_build_267) < 0 {
		return nil, ECGO_ERROR_SERVER_VERSION.throw()
	}

	dm_build_828.dm_build_808.ServerVersion = dm_build_832
	dm_build_828.dm_build_808.Malini2 = dm_build_816(dm_build_832, Dm_build_268) > 0
	dm_build_828.dm_build_808.Execute2 = dm_build_816(dm_build_832, Dm_build_269) > 0
	dm_build_828.dm_build_808.LobEmptyCompOrcl = dm_build_816(dm_build_832, Dm_build_270) > 0

	if dm_build_828.dm_build_405.dm_build_6.dmConnector.uKeyName != "" {
		dm_build_828.dm_build_812 = 1
	} else {
		dm_build_828.dm_build_812 = 0
	}

	if dm_build_828.dm_build_812 == 1 {
		dm_build_828.dm_build_405.dm_build_11 = dm_build_828.dm_build_405.dm_build_5.Dm_build_1381(16, dm_build_828.dm_build_808.getServerEncoding(), dm_build_828.dm_build_405.dm_build_6)
	}

	dm_build_828.dm_build_809 = -1
	dm_build_833 := false
	dm_build_834 := false
	dm_build_828.Dm_build_810 = -1
	if dm_build_830 > 0 {
		dm_build_828.dm_build_809 = int(dm_build_828.dm_build_405.dm_build_5.Dm_build_1344())
	}

	if dm_build_829 > 0 {

		if dm_build_828.dm_build_809 == -1 {
			dm_build_833 = true
		} else {
			dm_build_834 = true
		}

		dm_build_828.Dm_build_811 = dm_build_828.dm_build_405.dm_build_5.Dm_build_1369()
	}

	if dm_build_830 == 2 {
		dm_build_828.Dm_build_810 = dm_build_828.dm_build_405.dm_build_5.Dm_build_1344()
	}
	dm_build_828.dm_build_405.dm_build_8 = dm_build_833
	dm_build_828.dm_build_405.dm_build_9 = dm_build_834

	return nil, nil
}

type dm_build_835 struct {
	dm_build_404
}

func dm_build_836(dm_build_837 *dm_build_2, dm_build_838 *DmStatement) *dm_build_835 {
	dm_build_839 := new(dm_build_835)
	dm_build_839.dm_build_413(dm_build_837, Dm_build_275, dm_build_838)
	return dm_build_839
}

func (dm_build_841 *dm_build_835) dm_build_391() error {

	dm_build_841.dm_build_405.dm_build_5.Dm_build_1402(Dm_build_301, 1)
	return nil
}

func (dm_build_843 *dm_build_835) dm_build_395() (interface{}, error) {

	dm_build_843.dm_build_408.id = dm_build_843.dm_build_405.dm_build_5.Dm_build_1488(Dm_build_302)

	dm_build_843.dm_build_408.readBaseColName = dm_build_843.dm_build_405.dm_build_5.Dm_build_1482(Dm_build_301) == 1
	return nil, nil
}

type dm_build_844 struct {
	dm_build_404
	dm_build_845 int32
}

func dm_build_846(dm_build_847 *dm_build_2, dm_build_848 int32) *dm_build_844 {
	dm_build_849 := new(dm_build_844)
	dm_build_849.dm_build_409(dm_build_847, Dm_build_276)
	dm_build_849.dm_build_845 = dm_build_848
	return dm_build_849
}

func (dm_build_851 *dm_build_844) dm_build_392() {
	dm_build_851.dm_build_404.dm_build_392()
	dm_build_851.dm_build_405.dm_build_5.Dm_build_1410(Dm_build_302, dm_build_851.dm_build_845)
}

type dm_build_852 struct {
	dm_build_404
	dm_build_853 []uint32
}

func dm_build_854(dm_build_855 *dm_build_2, dm_build_856 []uint32) *dm_build_852 {
	dm_build_857 := new(dm_build_852)
	dm_build_857.dm_build_409(dm_build_855, Dm_build_296)
	dm_build_857.dm_build_853 = dm_build_856
	return dm_build_857
}

func (dm_build_859 *dm_build_852) dm_build_391() error {

	dm_build_859.dm_build_405.dm_build_5.Dm_build_1430(Dm_build_301, uint16(len(dm_build_859.dm_build_853)))

	for _, tableID := range dm_build_859.dm_build_853 {
		dm_build_859.dm_build_405.dm_build_5.Dm_build_1282(uint32(tableID))
	}

	return nil
}

func (dm_build_861 *dm_build_852) dm_build_395() (interface{}, error) {
	dm_build_862 := dm_build_861.dm_build_405.dm_build_5.Dm_build_1503(Dm_build_301)
	if dm_build_862 <= 0 {
		return nil, nil
	}

	dm_build_863 := make([]int64, dm_build_862)
	for i := 0; i < int(dm_build_862); i++ {
		dm_build_863[i] = dm_build_861.dm_build_405.dm_build_5.Dm_build_1347()
	}

	return dm_build_863, nil
}
