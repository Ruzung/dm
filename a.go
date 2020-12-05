/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"crypto/tls"
	"dm/security"
	"net"
	"strconv"
	"time"
)

const (
	Dm_build_0 = 8192
	Dm_build_1 = 2 * time.Second
)

type dm_build_2 struct {
	dm_build_3  *net.TCPConn
	dm_build_4  *tls.Conn
	dm_build_5  *Dm_build_1219
	dm_build_6  *DmConnection
	dm_build_7  security.Cipher
	dm_build_8  bool
	dm_build_9  bool
	dm_build_10 *security.DhKey
	dm_build_11 string
	dm_build_12 bool
}

func dm_build_13(dm_build_14 *DmConnection) (*dm_build_2, error) {
	dm_build_15, dm_build_16 := dm_build_18(dm_build_14.dmConnector.host+":"+strconv.Itoa(dm_build_14.dmConnector.port), time.Duration(dm_build_14.dmConnector.socketTimeout)*time.Second)
	if dm_build_16 != nil {
		return nil, dm_build_16
	}

	dm_build_17 := dm_build_2{}
	dm_build_17.dm_build_3 = dm_build_15
	dm_build_17.dm_build_5 = Dm_build_1222(Dm_build_272)
	dm_build_17.dm_build_6 = dm_build_14
	dm_build_17.dm_build_8 = false
	dm_build_17.dm_build_9 = false
	dm_build_17.dm_build_11 = ""
	dm_build_17.dm_build_12 = false
	dm_build_14.Access = &dm_build_17

	return &dm_build_17, nil
}

func dm_build_18(dm_build_19 string, dm_build_20 time.Duration) (*net.TCPConn, error) {
	dm_build_21, dm_build_22 := net.DialTimeout("tcp", dm_build_19, dm_build_20)
	if dm_build_22 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_19).throw()
	}

	if tcpConn, ok := dm_build_21.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_1)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_24 *dm_build_2) dm_build_23(dm_build_25 dm_build_389) bool {
	var dm_build_26 = dm_build_24.dm_build_6.dmConnector.compress
	if dm_build_25.dm_build_403() == Dm_build_299 || dm_build_26 == Dm_build_347 {
		return false
	}

	if dm_build_26 == Dm_build_345 {
		return true
	} else if dm_build_26 == Dm_build_346 {
		return !dm_build_24.dm_build_6.Local && dm_build_25.dm_build_401() > Dm_build_344
	}

	return false
}

func (dm_build_28 *dm_build_2) dm_build_27(dm_build_29 dm_build_389) bool {
	var dm_build_30 = dm_build_28.dm_build_6.dmConnector.compress
	if dm_build_29.dm_build_403() == Dm_build_299 || dm_build_30 == Dm_build_347 {
		return false
	}

	if dm_build_30 == Dm_build_345 {
		return true
	} else if dm_build_30 == Dm_build_346 {
		return dm_build_28.dm_build_5.Dm_build_1482(Dm_build_307) == 1
	}

	return false
}

func (dm_build_32 *dm_build_2) dm_build_31(dm_build_33 dm_build_389) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_35 := dm_build_33.dm_build_401()

	if dm_build_35 > 0 {

		if dm_build_32.dm_build_23(dm_build_33) {
			var retBytes, err = Compress(dm_build_32.dm_build_5, Dm_build_300, int(dm_build_35), int(dm_build_32.dm_build_6.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_32.dm_build_5.Dm_build_1233(Dm_build_300)

			dm_build_32.dm_build_5.Dm_build_1270(dm_build_35)

			dm_build_32.dm_build_5.Dm_build_1298(retBytes)

			dm_build_33.dm_build_402(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_32.dm_build_5.Dm_build_1402(Dm_build_307, 1)
		}

		if dm_build_32.dm_build_9 {
			dm_build_35 = dm_build_33.dm_build_401()
			var retBytes = dm_build_32.dm_build_7.Encrypt(dm_build_32.dm_build_5.Dm_build_1509(Dm_build_300, int(dm_build_35)), true)

			dm_build_32.dm_build_5.Dm_build_1233(Dm_build_300)

			dm_build_32.dm_build_5.Dm_build_1298(retBytes)

			dm_build_33.dm_build_402(int32(len(retBytes)))
		}
	}

	dm_build_33.dm_build_398()
	if dm_build_32.dm_build_263(dm_build_33) {
		if dm_build_32.dm_build_4 != nil {
			dm_build_32.dm_build_5.Dm_build_1236(0)
			dm_build_32.dm_build_5.Dm_build_1255(dm_build_32.dm_build_4)
		}
	} else {
		dm_build_32.dm_build_5.Dm_build_1236(0)
		dm_build_32.dm_build_5.Dm_build_1255(dm_build_32.dm_build_3)
	}
	return nil
}

func (dm_build_37 *dm_build_2) dm_build_36(dm_build_38 dm_build_389) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_40 := int32(0)
	if dm_build_37.dm_build_263(dm_build_38) {
		if dm_build_37.dm_build_4 != nil {
			dm_build_37.dm_build_5.Dm_build_1233(0)
			dm_build_37.dm_build_5.Dm_build_1249(dm_build_37.dm_build_4, Dm_build_300)
			dm_build_40 = dm_build_38.dm_build_401()
			if dm_build_40 > 0 {
				dm_build_37.dm_build_5.Dm_build_1249(dm_build_37.dm_build_4, int(dm_build_40))
			}
		}
	} else {

		dm_build_37.dm_build_5.Dm_build_1233(0)
		dm_build_37.dm_build_5.Dm_build_1249(dm_build_37.dm_build_3, Dm_build_300)
		dm_build_40 = dm_build_38.dm_build_401()

		if dm_build_40 > 0 {
			dm_build_37.dm_build_5.Dm_build_1249(dm_build_37.dm_build_3, int(dm_build_40))
		}
	}

	dm_build_38.dm_build_399()

	if dm_build_40 <= 0 {
		return nil
	}

	if dm_build_37.dm_build_9 {
		dm_build_40 = dm_build_38.dm_build_401()
		ebytes := dm_build_37.dm_build_5.Dm_build_1509(Dm_build_300, int(dm_build_40))
		bytes, err := dm_build_37.dm_build_7.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_37.dm_build_5.Dm_build_1233(Dm_build_300)
		dm_build_37.dm_build_5.Dm_build_1298(bytes)
		dm_build_38.dm_build_402(int32(len(bytes)))
	}

	if dm_build_37.dm_build_27(dm_build_38) {

		dm_build_40 = dm_build_38.dm_build_401()
		cbytes := dm_build_37.dm_build_5.Dm_build_1509(Dm_build_300+ULINT_SIZE, int(dm_build_40-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_37.dm_build_6.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_37.dm_build_5.Dm_build_1233(Dm_build_300)
		dm_build_37.dm_build_5.Dm_build_1298(bytes)
		dm_build_38.dm_build_402(int32(len(bytes)))
	}
	return nil
}

func (dm_build_42 *dm_build_2) dm_build_41(dm_build_43 dm_build_389) (dm_build_44 interface{}, dm_build_45 error) {
	dm_build_45 = dm_build_43.dm_build_393(dm_build_43)
	if dm_build_45 != nil {
		return nil, dm_build_45
	}

	dm_build_45 = dm_build_42.dm_build_31(dm_build_43)
	if dm_build_45 != nil {
		return nil, dm_build_45
	}

	dm_build_45 = dm_build_42.dm_build_36(dm_build_43)
	if dm_build_45 != nil {
		return nil, dm_build_45
	}

	return dm_build_43.dm_build_397(dm_build_43)
}

func (dm_build_47 *dm_build_2) dm_build_46() (*dm_build_807, error) {

	Dm_build_48 := dm_build_813(dm_build_47)
	_, dm_build_49 := dm_build_47.dm_build_41(Dm_build_48)
	if dm_build_49 != nil {
		return nil, dm_build_49
	}

	return Dm_build_48, nil
}

func (dm_build_51 *dm_build_2) dm_build_50() error {

	dm_build_52 := dm_build_684(dm_build_51)
	_, dm_build_53 := dm_build_51.dm_build_41(dm_build_52)
	if dm_build_53 != nil {
		return dm_build_53
	}

	return nil
}

func (dm_build_55 *dm_build_2) dm_build_54() error {

	var dm_build_56 *dm_build_807
	var err error
	if dm_build_56, err = dm_build_55.dm_build_46(); err != nil {
		return err
	}

	if dm_build_55.dm_build_6.sslEncrypt == 2 {
		if err = dm_build_55.dm_build_259(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_55.dm_build_6.sslEncrypt == 1 {
		if err = dm_build_55.dm_build_259(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_55.dm_build_9 || dm_build_55.dm_build_8 {
		k, err := dm_build_55.dm_build_249()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_56.Dm_build_811)
		encryptType := dm_build_56.dm_build_809
		hashType := int(dm_build_56.Dm_build_810)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_55.dm_build_252(encryptType, sessionKey, dm_build_55.dm_build_6.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_55.dm_build_50(); err != nil {
		return err
	}
	return nil
}

func (dm_build_59 *dm_build_2) Dm_build_58(dm_build_60 *DmStatement) error {
	dm_build_61 := dm_build_836(dm_build_59, dm_build_60)
	_, dm_build_62 := dm_build_59.dm_build_41(dm_build_61)
	if dm_build_62 != nil {
		return dm_build_62
	}

	return nil
}

func (dm_build_64 *dm_build_2) Dm_build_63(dm_build_65 int32) error {
	dm_build_66 := dm_build_846(dm_build_64, dm_build_65)
	_, dm_build_67 := dm_build_64.dm_build_41(dm_build_66)
	if dm_build_67 != nil {
		return dm_build_67
	}

	return nil
}

func (dm_build_69 *dm_build_2) Dm_build_68(dm_build_70 *DmStatement, dm_build_71 bool, dm_build_72 int16) (*execInfo, error) {
	dm_build_73 := dm_build_720(dm_build_69, dm_build_70, dm_build_71, dm_build_72)
	dm_build_74, dm_build_75 := dm_build_69.dm_build_41(dm_build_73)
	if dm_build_75 != nil {
		return nil, dm_build_75
	}
	return dm_build_74.(*execInfo), nil
}

func (dm_build_77 *dm_build_2) Dm_build_76(dm_build_78 *DmStatement, dm_build_79 int16) (*execInfo, error) {
	return dm_build_77.Dm_build_68(dm_build_78, false, Dm_build_351)
}

func (dm_build_81 *dm_build_2) Dm_build_80(dm_build_82 *DmStatement, dm_build_83 []OptParameter) (*execInfo, error) {
	dm_build_84, dm_build_85 := dm_build_81.dm_build_41(dm_build_482(dm_build_81, dm_build_82, dm_build_83))
	if dm_build_85 != nil {
		return nil, dm_build_85
	}

	return dm_build_84.(*execInfo), nil
}

func (dm_build_87 *dm_build_2) Dm_build_86(dm_build_88 *DmStatement, dm_build_89 int16) (*execInfo, error) {
	return dm_build_87.Dm_build_68(dm_build_88, true, dm_build_89)
}

func (dm_build_91 *dm_build_2) Dm_build_90(dm_build_92 *DmStatement, dm_build_93 [][]interface{}) (*execInfo, error) {
	dm_build_94 := dm_build_505(dm_build_91, dm_build_92, dm_build_93)
	dm_build_95, dm_build_96 := dm_build_91.dm_build_41(dm_build_94)
	if dm_build_96 != nil {
		return nil, dm_build_96
	}
	return dm_build_95.(*execInfo), nil
}

func (dm_build_98 *dm_build_2) Dm_build_97(dm_build_99 *DmStatement, dm_build_100 [][]interface{}) (*execInfo, error) {
	var dm_build_101, dm_build_102 = 0, 0
	var dm_build_103 = len(dm_build_100)
	var dm_build_104 [][]interface{}
	var dm_build_105 = NewExceInfo()
	dm_build_105.updateCounts = make([]int64, dm_build_103)
	var dm_build_106 = false
	var dm_build_107 = false
	for dm_build_101 < dm_build_103 {
		for dm_build_102 = dm_build_101; dm_build_102 < dm_build_103; dm_build_102++ {
			paramData := dm_build_100[dm_build_102]
			bindData := make([]interface{}, dm_build_99.paramCount)
			dm_build_106 = false
			for icol := 0; icol < int(dm_build_99.paramCount); icol++ {
				if dm_build_99.params[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_98.dm_build_232(bindData, paramData, icol) {
					dm_build_106 = true
					break
				}
			}

			if dm_build_106 {
				break
			}
			dm_build_104 = append(dm_build_104, bindData)
		}

		if dm_build_102 != dm_build_101 {
			tmpExecInfo, err := dm_build_98.Dm_build_90(dm_build_99, dm_build_104)
			if err != nil {
				return nil, err
			}
			dm_build_104 = dm_build_104[0:0]
			if dm_build_102-dm_build_101 == 1 {
				dm_build_105.updateCounts[dm_build_101] = tmpExecInfo.updateCount
			} else if tmpExecInfo.updateCounts != nil {
				copy(dm_build_105.updateCounts[dm_build_101:dm_build_102], tmpExecInfo.updateCounts[0:dm_build_102-dm_build_101])
			}
		}

		if dm_build_102 < dm_build_103 {
			tmpExecInfo, err := dm_build_98.Dm_build_108(dm_build_99, dm_build_100[dm_build_102], dm_build_107)
			if err != nil {
				return nil, err
			}

			dm_build_107 = true
			dm_build_105.updateCounts[dm_build_102] = tmpExecInfo.updateCount
		}

		dm_build_101 = dm_build_102 + 1
	}
	for _, i := range dm_build_105.updateCounts {
		dm_build_105.updateCount += i
	}
	return dm_build_105, nil
}

func (dm_build_109 *dm_build_2) Dm_build_108(dm_build_110 *DmStatement, dm_build_111 []interface{}, dm_build_112 bool) (*execInfo, error) {

	var dm_build_113 = make([]interface{}, dm_build_110.paramCount)
	for icol := 0; icol < int(dm_build_110.paramCount); icol++ {
		if dm_build_110.params[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_109.dm_build_232(dm_build_113, dm_build_111, icol) {

			if !dm_build_112 {
				preExecute := dm_build_710(dm_build_109, dm_build_110, dm_build_110.params)
				dm_build_109.dm_build_41(preExecute)
				dm_build_112 = true
			}

			dm_build_109.dm_build_238(dm_build_110, dm_build_110.params[icol], icol, dm_build_111[icol].(iOffRowBinder))
			dm_build_113[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_114 = make([][]interface{}, 1, 1)
	dm_build_114[0] = dm_build_113

	dm_build_115 := dm_build_505(dm_build_109, dm_build_110, dm_build_114)
	dm_build_116, dm_build_117 := dm_build_109.dm_build_41(dm_build_115)
	if dm_build_117 != nil {
		return nil, dm_build_117
	}
	return dm_build_116.(*execInfo), nil
}

func (dm_build_119 *dm_build_2) Dm_build_118(dm_build_120 *DmStatement, dm_build_121 int16) (*execInfo, error) {
	dm_build_122 := dm_build_698(dm_build_119, dm_build_120, dm_build_121)

	dm_build_123, dm_build_124 := dm_build_119.dm_build_41(dm_build_122)
	if dm_build_124 != nil {
		return nil, dm_build_124
	}
	return dm_build_123.(*execInfo), nil
}

func (dm_build_126 *dm_build_2) Dm_build_125(dm_build_127 *innerRows, dm_build_128 int64) (*execInfo, error) {
	dm_build_129 := dm_build_605(dm_build_126, dm_build_127, dm_build_128, INT64_MAX)
	dm_build_130, dm_build_131 := dm_build_126.dm_build_41(dm_build_129)
	if dm_build_131 != nil {
		return nil, dm_build_131
	}
	return dm_build_130.(*execInfo), nil
}

func (dm_build_133 *dm_build_2) Commit() error {
	dm_build_134 := dm_build_468(dm_build_133)
	_, dm_build_135 := dm_build_133.dm_build_41(dm_build_134)
	if dm_build_135 != nil {
		return dm_build_135
	}

	return nil
}

func (dm_build_137 *dm_build_2) Rollback() error {
	dm_build_138 := dm_build_758(dm_build_137)
	_, dm_build_139 := dm_build_137.dm_build_41(dm_build_138)
	if dm_build_139 != nil {
		return dm_build_139
	}

	return nil
}

func (dm_build_141 *dm_build_2) Dm_build_140(dm_build_142 *DmConnection) error {
	dm_build_143 := dm_build_763(dm_build_141, dm_build_142.IsoLevel)
	_, dm_build_144 := dm_build_141.dm_build_41(dm_build_143)
	if dm_build_144 != nil {
		return dm_build_144
	}

	return nil
}

func (dm_build_146 *dm_build_2) Dm_build_145(dm_build_147 *DmStatement, dm_build_148 string) error {
	dm_build_149 := dm_build_473(dm_build_146, dm_build_147, dm_build_148)
	_, dm_build_150 := dm_build_146.dm_build_41(dm_build_149)
	if dm_build_150 != nil {
		return dm_build_150
	}

	return nil
}

func (dm_build_152 *dm_build_2) Dm_build_151(dm_build_153 []uint32) ([]int64, error) {
	dm_build_154 := dm_build_854(dm_build_152, dm_build_153)
	dm_build_155, dm_build_156 := dm_build_152.dm_build_41(dm_build_154)
	if dm_build_156 != nil {
		return nil, dm_build_156
	}
	return dm_build_155.([]int64), nil
}

func (dm_build_158 *dm_build_2) Close() error {
	if dm_build_158.dm_build_12 {
		return nil
	}

	dm_build_159 := dm_build_158.dm_build_3.Close()
	if dm_build_159 != nil {
		return dm_build_159
	}

	dm_build_158.dm_build_6 = nil
	dm_build_158.dm_build_12 = true
	return nil
}

func (dm_build_161 *dm_build_2) dm_build_160(dm_build_162 *lob) (int64, error) {
	dm_build_163 := dm_build_636(dm_build_161, dm_build_162)
	dm_build_164, dm_build_165 := dm_build_161.dm_build_41(dm_build_163)
	if dm_build_165 != nil {
		return 0, dm_build_165
	}
	return dm_build_164.(int64), nil
}

func (dm_build_167 *dm_build_2) dm_build_166(dm_build_168 *lob, dm_build_169 int32, dm_build_170 int32) ([]byte, error) {
	dm_build_171 := dm_build_623(dm_build_167, dm_build_168, int(dm_build_169), int(dm_build_170))
	dm_build_172, dm_build_173 := dm_build_167.dm_build_41(dm_build_171)
	if dm_build_173 != nil {
		return nil, dm_build_173
	}
	return dm_build_172.([]byte), nil
}

func (dm_build_175 *dm_build_2) dm_build_174(dm_build_176 *DmBlob, dm_build_177 int32, dm_build_178 int32) ([]byte, error) {
	var dm_build_179 = make([]byte, dm_build_178)
	var dm_build_180 int32 = 0
	var dm_build_181 int32 = 0
	var dm_build_182 []byte
	var dm_build_183 error
	for dm_build_180 < dm_build_178 {
		dm_build_181 = dm_build_178 - dm_build_180
		if dm_build_181 > Dm_build_384 {
			dm_build_181 = Dm_build_384
		}
		dm_build_182, dm_build_183 = dm_build_175.dm_build_166(&dm_build_176.lob, dm_build_177, dm_build_178)
		if dm_build_183 != nil {
			return nil, dm_build_183
		}
		if dm_build_182 == nil || len(dm_build_182) == 0 {
			break
		}
		Dm_build_865.Dm_build_921(dm_build_179, int(dm_build_180), dm_build_182, 0, len(dm_build_182))
		dm_build_180 += int32(len(dm_build_182))
		dm_build_177 += int32(len(dm_build_182))
		if dm_build_176.readOver {
			break
		}
	}
	return dm_build_179, nil
}

func (dm_build_185 *dm_build_2) dm_build_184(dm_build_186 *DmClob, dm_build_187 int32, dm_build_188 int32) (string, error) {
	var dm_build_189 = ""
	var dm_build_190 int32 = 0
	var dm_build_191 int32 = 0
	var dm_build_192 []byte
	var dm_build_193 string
	var dm_build_194 error
	for dm_build_190 < dm_build_188 {
		dm_build_191 = dm_build_188 - dm_build_190
		if dm_build_191 > Dm_build_384/2 {
			dm_build_191 = Dm_build_384 / 2
		}
		dm_build_192, dm_build_194 = dm_build_185.dm_build_166(&dm_build_186.lob, dm_build_187, dm_build_188)
		if dm_build_194 != nil {
			return "", dm_build_194
		}
		if dm_build_192 == nil || len(dm_build_192) == 0 {
			break
		}
		dm_build_193 = Dm_build_865.Dm_build_1019(dm_build_192, 0, len(dm_build_192), dm_build_186.serverEncoding, dm_build_185.dm_build_6)
		dm_build_189 += dm_build_193
		dm_build_190 += int32(len(dm_build_193))
		dm_build_187 += int32(len(dm_build_192))
		if dm_build_186.readOver {
			break
		}
	}
	return dm_build_189, nil
}

func (dm_build_196 *dm_build_2) dm_build_195(dm_build_197 *DmClob, dm_build_198 int, dm_build_199 string, dm_build_200 string) (int, error) {
	var dm_build_201 = Dm_build_865.Dm_build_1075(dm_build_199, dm_build_200, dm_build_196.dm_build_6)
	var dm_build_202 = 0
	var dm_build_203 = len(dm_build_201)
	var dm_build_204 = 0
	var dm_build_205 = 0
	var dm_build_206 = 0
	var dm_build_207 = dm_build_203/Dm_build_383 + 1
	var dm_build_208 byte = 0
	var dm_build_209 byte = 0x01
	var dm_build_210 byte = 0x02
	for i := 0; i < dm_build_207; i++ {
		dm_build_208 = 0
		if i == 0 {
			dm_build_208 |= dm_build_209
		}
		if i == dm_build_207-1 {
			dm_build_208 |= dm_build_210
		}
		dm_build_206 = dm_build_203 - dm_build_205
		if dm_build_206 > Dm_build_383 {
			dm_build_206 = Dm_build_383
		}

		setLobData := dm_build_777(dm_build_196, &dm_build_197.lob, dm_build_208, dm_build_198, dm_build_201, dm_build_202, dm_build_206)
		ret, err := dm_build_196.dm_build_41(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_204, nil
		} else {
			dm_build_198 += int(tmp)
			dm_build_204 += int(tmp)
			dm_build_205 += dm_build_206
			dm_build_202 += dm_build_206
		}
	}
	return dm_build_204, nil
}

func (dm_build_212 *dm_build_2) dm_build_211(dm_build_213 *DmBlob, dm_build_214 int, dm_build_215 []byte) (int, error) {
	var dm_build_216 = 0
	var dm_build_217 = len(dm_build_215)
	var dm_build_218 = 0
	var dm_build_219 = 0
	var dm_build_220 = 0
	var dm_build_221 = dm_build_217/Dm_build_383 + 1
	var dm_build_222 byte = 0
	var dm_build_223 byte = 0x01
	var dm_build_224 byte = 0x02
	for i := 0; i < dm_build_221; i++ {
		dm_build_222 = 0
		if i == 0 {
			dm_build_222 |= dm_build_223
		}
		if i == dm_build_221-1 {
			dm_build_222 |= dm_build_224
		}
		dm_build_220 = dm_build_217 - dm_build_219
		if dm_build_220 > Dm_build_383 {
			dm_build_220 = Dm_build_383
		}

		setLobData := dm_build_777(dm_build_212, &dm_build_213.lob, dm_build_222, dm_build_214, dm_build_215, dm_build_216, dm_build_220)
		ret, err := dm_build_212.dm_build_41(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_218, nil
		} else {
			dm_build_214 += int(tmp)
			dm_build_218 += int(tmp)
			dm_build_219 += dm_build_220
			dm_build_216 += dm_build_220
		}
	}
	return dm_build_218, nil
}

func (dm_build_226 *dm_build_2) dm_build_225(dm_build_227 *lob, dm_build_228 int) (int64, error) {
	dm_build_229 := dm_build_647(dm_build_226, dm_build_227, dm_build_228)
	dm_build_230, dm_build_231 := dm_build_226.dm_build_41(dm_build_229)
	if dm_build_231 != nil {
		return dm_build_227.length, dm_build_231
	}
	return dm_build_230.(int64), nil
}

func (dm_build_233 *dm_build_2) dm_build_232(dm_build_234 []interface{}, dm_build_235 []interface{}, dm_build_236 int) bool {
	var dm_build_237 = false
	if dm_build_236 >= len(dm_build_235) || dm_build_235[dm_build_236] == nil {
		dm_build_234[dm_build_236] = ParamDataEnum_Null
	} else if binder, ok := dm_build_235[dm_build_236].(iOffRowBinder); ok {
		dm_build_237 = true
		dm_build_234[dm_build_236] = ParamDataEnum_OFF_ROW
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_233.dm_build_6) {
			dm_build_234[dm_build_236] = &lobCtl{lob.buildCtlData()}
			dm_build_237 = false
		}
	} else {
		dm_build_234[dm_build_236] = dm_build_235[dm_build_236]
	}
	return dm_build_237
}

func (dm_build_239 *dm_build_2) dm_build_238(dm_build_240 *DmStatement, dm_build_241 parameter, dm_build_242 int, dm_build_243 iOffRowBinder) error {
	var dm_build_244 = Dm_build_1145()
	dm_build_243.read(dm_build_244)
	var dm_build_245 = 0
	for !dm_build_243.isReadOver() || dm_build_244.Dm_build_1146() > 0 {
		if !dm_build_243.isReadOver() && dm_build_244.Dm_build_1146() < Dm_build_383 {
			dm_build_243.read(dm_build_244)
		}
		if dm_build_244.Dm_build_1146() > Dm_build_383 {
			dm_build_245 = Dm_build_383
		} else {
			dm_build_245 = dm_build_244.Dm_build_1146()
		}

		putData := dm_build_748(dm_build_239, dm_build_240, int16(dm_build_242), dm_build_244, int32(dm_build_245))
		_, err := dm_build_239.dm_build_41(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_247 *dm_build_2) dm_build_246() ([]byte, error) {
	var dm_build_248 error
	if dm_build_247.dm_build_10 == nil {
		if dm_build_247.dm_build_10, dm_build_248 = security.NewClientKeyPair(); dm_build_248 != nil {
			return nil, dm_build_248
		}
	}
	return security.Bn2Bytes(dm_build_247.dm_build_10.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_250 *dm_build_2) dm_build_249() (*security.DhKey, error) {
	var dm_build_251 error
	if dm_build_250.dm_build_10 == nil {
		if dm_build_250.dm_build_10, dm_build_251 = security.NewClientKeyPair(); dm_build_251 != nil {
			return nil, dm_build_251
		}
	}
	return dm_build_250.dm_build_10, nil
}

func (dm_build_253 *dm_build_2) dm_build_252(dm_build_254 int, dm_build_255 []byte, dm_build_256 string, dm_build_257 int) (dm_build_258 error) {
	if dm_build_254 > 0 && dm_build_254 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_255 != nil {
		dm_build_253.dm_build_7, dm_build_258 = security.NewSymmCipher(dm_build_254, dm_build_255)
	} else if dm_build_254 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_253.dm_build_7, dm_build_258 = security.NewThirdPartCipher(dm_build_254, dm_build_255, dm_build_256, dm_build_257); dm_build_258 != nil {
			dm_build_258 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_258.Error()).throw()
		}
	}
	return
}

func (dm_build_260 *dm_build_2) dm_build_259(dm_build_261 bool) (dm_build_262 error) {
	if dm_build_260.dm_build_4, dm_build_262 = security.NewTLSFromTCP(dm_build_260.dm_build_3, dm_build_260.dm_build_6.dmConnector.sslCertPath, dm_build_260.dm_build_6.dmConnector.sslKeyPath, dm_build_260.dm_build_6.dmConnector.user); dm_build_262 != nil {
		return
	}
	if !dm_build_261 {
		dm_build_260.dm_build_4 = nil
	}
	return
}

func (dm_build_264 *dm_build_2) dm_build_263(dm_build_265 dm_build_389) bool {
	return dm_build_265.dm_build_403() != Dm_build_299 && dm_build_264.dm_build_6.sslEncrypt == 1
}
