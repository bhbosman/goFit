package fitDecl

import strconv "strconv"

type GarminProduct uint16

const (
	GarminProduct_Hrm1                       GarminProduct = 1
	GarminProduct_Axh01                      GarminProduct = 2
	GarminProduct_Axb01                      GarminProduct = 3
	GarminProduct_Axb02                      GarminProduct = 4
	GarminProduct_Hrm2ss                     GarminProduct = 5
	GarminProduct_DsiAlf02                   GarminProduct = 6
	GarminProduct_Hrm3ss                     GarminProduct = 7
	GarminProduct_HrmRunSingleByteProductId  GarminProduct = 8
	GarminProduct_Bsm                        GarminProduct = 9
	GarminProduct_Bcm                        GarminProduct = 10
	GarminProduct_Axs01                      GarminProduct = 11
	GarminProduct_HrmTriSingleByteProductId  GarminProduct = 12
	GarminProduct_Hrm4RunSingleByteProductId GarminProduct = 13
	GarminProduct_Fr225SingleByteProductId   GarminProduct = 14
	GarminProduct_Gen3BsmSingleByteProductId GarminProduct = 15
	GarminProduct_Gen3BcmSingleByteProductId GarminProduct = 16
	GarminProduct_HrmFitSingleByteProductId  GarminProduct = 22
	GarminProduct_OHR                        GarminProduct = 255
	GarminProduct_Fr301China                 GarminProduct = 473
	GarminProduct_Fr301Japan                 GarminProduct = 474
	GarminProduct_Fr301Korea                 GarminProduct = 475
	GarminProduct_Fr301Taiwan                GarminProduct = 494
	GarminProduct_Fr405                      GarminProduct = 717
	GarminProduct_Fr50                       GarminProduct = 782
	GarminProduct_Fr405Japan                 GarminProduct = 987
	GarminProduct_Fr60                       GarminProduct = 988
	GarminProduct_DsiAlf01                   GarminProduct = 1011
	GarminProduct_Fr310xt                    GarminProduct = 1018
	GarminProduct_Edge500                    GarminProduct = 1036
	GarminProduct_Fr110                      GarminProduct = 1124
	GarminProduct_Edge800                    GarminProduct = 1169
	GarminProduct_Edge500Taiwan              GarminProduct = 1199
	GarminProduct_Edge500Japan               GarminProduct = 1213
	GarminProduct_Chirp                      GarminProduct = 1253
	GarminProduct_Fr110Japan                 GarminProduct = 1274
	GarminProduct_Edge200                    GarminProduct = 1325
	GarminProduct_Fr910xt                    GarminProduct = 1328
	GarminProduct_Edge800Taiwan              GarminProduct = 1333
	GarminProduct_Edge800Japan               GarminProduct = 1334
	GarminProduct_Alf04                      GarminProduct = 1341
	GarminProduct_Fr610                      GarminProduct = 1345
	GarminProduct_Fr210Japan                 GarminProduct = 1360
	GarminProduct_VectorSs                   GarminProduct = 1380
	GarminProduct_VectorCp                   GarminProduct = 1381
	GarminProduct_Edge800China               GarminProduct = 1386
	GarminProduct_Edge500China               GarminProduct = 1387
	GarminProduct_ApproachG10                GarminProduct = 1405
	GarminProduct_Fr610Japan                 GarminProduct = 1410
	GarminProduct_Edge500Korea               GarminProduct = 1422
	GarminProduct_Fr70                       GarminProduct = 1436
	GarminProduct_Fr310xt4t                  GarminProduct = 1446
	GarminProduct_Amx                        GarminProduct = 1461
	GarminProduct_Fr10                       GarminProduct = 1482
	GarminProduct_Edge800Korea               GarminProduct = 1497
	GarminProduct_Swim                       GarminProduct = 1499
	GarminProduct_Fr910xtChina               GarminProduct = 1537
	GarminProduct_Fenix                      GarminProduct = 1551
	GarminProduct_Edge200Taiwan              GarminProduct = 1555
	GarminProduct_Edge510                    GarminProduct = 1561
	GarminProduct_Edge810                    GarminProduct = 1567
	GarminProduct_Tempe                      GarminProduct = 1570
	GarminProduct_Fr910xtJapan               GarminProduct = 1600
	GarminProduct_Fr620                      GarminProduct = 1623
	GarminProduct_Fr220                      GarminProduct = 1632
	GarminProduct_Fr910xtKorea               GarminProduct = 1664
	GarminProduct_Fr10Japan                  GarminProduct = 1688
	GarminProduct_Edge810Japan               GarminProduct = 1721
	GarminProduct_VirbElite                  GarminProduct = 1735
	GarminProduct_EdgeTouring                GarminProduct = 1736
	GarminProduct_Edge510Japan               GarminProduct = 1742
	GarminProduct_HrmTri                     GarminProduct = 1743
	GarminProduct_HrmRun                     GarminProduct = 1752
	GarminProduct_Fr920xt                    GarminProduct = 1765
	GarminProduct_Edge510Asia                GarminProduct = 1821
	GarminProduct_Edge810China               GarminProduct = 1822
	GarminProduct_Edge810Taiwan              GarminProduct = 1823
	GarminProduct_Edge1000                   GarminProduct = 1836
	GarminProduct_VivoFit                    GarminProduct = 1837
	GarminProduct_VirbRemote                 GarminProduct = 1853
	GarminProduct_VivoKi                     GarminProduct = 1885
	GarminProduct_Fr15                       GarminProduct = 1903
	GarminProduct_VivoActive                 GarminProduct = 1907
	GarminProduct_Edge510Korea               GarminProduct = 1918
	GarminProduct_Fr620Japan                 GarminProduct = 1928
	GarminProduct_Fr620China                 GarminProduct = 1929
	GarminProduct_Fr220Japan                 GarminProduct = 1930
	GarminProduct_Fr220China                 GarminProduct = 1931
	GarminProduct_ApproachS6                 GarminProduct = 1936
	GarminProduct_VivoSmart                  GarminProduct = 1956
	GarminProduct_Fenix2                     GarminProduct = 1967
	GarminProduct_Epix                       GarminProduct = 1988
	GarminProduct_Fenix3                     GarminProduct = 2050
	GarminProduct_Edge1000Taiwan             GarminProduct = 2052
	GarminProduct_Edge1000Japan              GarminProduct = 2053
	GarminProduct_Fr15Japan                  GarminProduct = 2061
	GarminProduct_Edge520                    GarminProduct = 2067
	GarminProduct_Edge1000China              GarminProduct = 2070
	GarminProduct_Fr620Russia                GarminProduct = 2072
	GarminProduct_Fr220Russia                GarminProduct = 2073
	GarminProduct_VectorS                    GarminProduct = 2079
	GarminProduct_Edge1000Korea              GarminProduct = 2100
	GarminProduct_Fr920xtTaiwan              GarminProduct = 2130
	GarminProduct_Fr920xtChina               GarminProduct = 2131
	GarminProduct_Fr920xtJapan               GarminProduct = 2132
	GarminProduct_Virbx                      GarminProduct = 2134
	GarminProduct_VivoSmartApac              GarminProduct = 2135
	GarminProduct_EtrexTouch                 GarminProduct = 2140
	GarminProduct_Edge25                     GarminProduct = 2147
	GarminProduct_Fr25                       GarminProduct = 2148
	GarminProduct_VivoFit2                   GarminProduct = 2150
	GarminProduct_Fr225                      GarminProduct = 2153
	GarminProduct_Fr630                      GarminProduct = 2156
	GarminProduct_Fr230                      GarminProduct = 2157
	GarminProduct_Fr735xt                    GarminProduct = 2158
	GarminProduct_VivoActiveApac             GarminProduct = 2160
	GarminProduct_Vector2                    GarminProduct = 2161
	GarminProduct_Vector2s                   GarminProduct = 2162
	GarminProduct_Virbxe                     GarminProduct = 2172
	GarminProduct_Fr620Taiwan                GarminProduct = 2173
	GarminProduct_Fr220Taiwan                GarminProduct = 2174
	GarminProduct_Truswing                   GarminProduct = 2175
	GarminProduct_D2airvenu                  GarminProduct = 2187
	GarminProduct_Fenix3China                GarminProduct = 2188
	GarminProduct_Fenix3Twn                  GarminProduct = 2189
	GarminProduct_VariaHeadlight             GarminProduct = 2192
	GarminProduct_VariaTaillightOld          GarminProduct = 2193
	GarminProduct_EdgeExplore1000            GarminProduct = 2204
	GarminProduct_Fr225Asia                  GarminProduct = 2219
	GarminProduct_VariaRadarTaillight        GarminProduct = 2225
	GarminProduct_VariaRadarDisplay          GarminProduct = 2226
	GarminProduct_Edge20                     GarminProduct = 2238
	GarminProduct_Edge520Asia                GarminProduct = 2260
	GarminProduct_Edge520Japan               GarminProduct = 2261
	GarminProduct_D2Bravo                    GarminProduct = 2262
	GarminProduct_ApproachS20                GarminProduct = 2266
	GarminProduct_VivoSmart2                 GarminProduct = 2271
	GarminProduct_Edge1000Thai               GarminProduct = 2274
	GarminProduct_VariaRemote                GarminProduct = 2276
	GarminProduct_Edge25Asia                 GarminProduct = 2288
	GarminProduct_Edge25Jpn                  GarminProduct = 2289
	GarminProduct_Edge20Asia                 GarminProduct = 2290
	GarminProduct_ApproachX40                GarminProduct = 2292
	GarminProduct_Fenix3Japan                GarminProduct = 2293
	GarminProduct_VivoSmartEmea              GarminProduct = 2294
	GarminProduct_Fr630Asia                  GarminProduct = 2310
	GarminProduct_Fr630Jpn                   GarminProduct = 2311
	GarminProduct_Fr230Jpn                   GarminProduct = 2313
	GarminProduct_Hrm4Run                    GarminProduct = 2327
	GarminProduct_EpixJapan                  GarminProduct = 2332
	GarminProduct_VivoActiveHr               GarminProduct = 2337
	GarminProduct_VivoSmartGpsHr             GarminProduct = 2347
	GarminProduct_VivoSmartHr                GarminProduct = 2348
	GarminProduct_VivoSmartHrAsia            GarminProduct = 2361
	GarminProduct_VivoSmartGpsHrAsia         GarminProduct = 2362
	GarminProduct_VivoMove                   GarminProduct = 2368
	GarminProduct_VariaTaillight             GarminProduct = 2379
	GarminProduct_Fr235Asia                  GarminProduct = 2396
	GarminProduct_Fr235Japan                 GarminProduct = 2397
	GarminProduct_VariaVision                GarminProduct = 2398
	GarminProduct_VivoFit3                   GarminProduct = 2406
	GarminProduct_Fenix3Korea                GarminProduct = 2407
	GarminProduct_Fenix3Sea                  GarminProduct = 2408
	GarminProduct_Fenix3Hr                   GarminProduct = 2413
	GarminProduct_VirbUltra30                GarminProduct = 2417
	GarminProduct_IndexSmartScale            GarminProduct = 2429
	GarminProduct_Fr235                      GarminProduct = 2431
	GarminProduct_Fenix3Chronos              GarminProduct = 2432
	GarminProduct_Oregon7xx                  GarminProduct = 2441
	GarminProduct_Rino7xx                    GarminProduct = 2444
	GarminProduct_EpixKorea                  GarminProduct = 2457
	GarminProduct_Fenix3HrChn                GarminProduct = 2473
	GarminProduct_Fenix3HrTwn                GarminProduct = 2474
	GarminProduct_Fenix3HrJpn                GarminProduct = 2475
	GarminProduct_Fenix3HrSea                GarminProduct = 2476
	GarminProduct_Fenix3HrKor                GarminProduct = 2477
	GarminProduct_Nautix                     GarminProduct = 2496
	GarminProduct_VivoActiveHrApac           GarminProduct = 2497
	GarminProduct_Fr35                       GarminProduct = 2503
	GarminProduct_Oregon7xxWw                GarminProduct = 2512
	GarminProduct_Edge820                    GarminProduct = 2530
	GarminProduct_EdgeExplore820             GarminProduct = 2531
	GarminProduct_Fr735xtApac                GarminProduct = 2533
	GarminProduct_Fr735xtJapan               GarminProduct = 2534
	GarminProduct_Fenix5s                    GarminProduct = 2544
	GarminProduct_D2BravoTitanium            GarminProduct = 2547
	GarminProduct_VariaUt800                 GarminProduct = 2567
	GarminProduct_RunningDynamicsPod         GarminProduct = 2593
	GarminProduct_Edge820China               GarminProduct = 2599
	GarminProduct_Edge820Japan               GarminProduct = 2600
	GarminProduct_Fenix5x                    GarminProduct = 2604
	GarminProduct_VivoFitJr                  GarminProduct = 2606
	GarminProduct_VivoSmart3                 GarminProduct = 2622
	GarminProduct_VivoSport                  GarminProduct = 2623
	GarminProduct_Edge820Taiwan              GarminProduct = 2628
	GarminProduct_Edge820Korea               GarminProduct = 2629
	GarminProduct_Edge820Sea                 GarminProduct = 2630
	GarminProduct_Fr35Hebrew                 GarminProduct = 2650
	GarminProduct_ApproachS60                GarminProduct = 2656
	GarminProduct_Fr35Apac                   GarminProduct = 2667
	GarminProduct_Fr35Japan                  GarminProduct = 2668
	GarminProduct_Fenix3ChronosAsia          GarminProduct = 2675
	GarminProduct_Virb360                    GarminProduct = 2687
	GarminProduct_Fr935                      GarminProduct = 2691
	GarminProduct_Fenix5                     GarminProduct = 2697
	GarminProduct_Vivoactive3                GarminProduct = 2700
	GarminProduct_Fr235ChinaNfc              GarminProduct = 2733
	GarminProduct_Foretrex601701             GarminProduct = 2769
	GarminProduct_VivoMoveHr                 GarminProduct = 2772
	GarminProduct_Edge1030                   GarminProduct = 2713
	GarminProduct_Fr35Sea                    GarminProduct = 2727
	GarminProduct_Vector3                    GarminProduct = 2787
	GarminProduct_Fenix5Asia                 GarminProduct = 2796
	GarminProduct_Fenix5sAsia                GarminProduct = 2797
	GarminProduct_Fenix5xAsia                GarminProduct = 2798
	GarminProduct_ApproachZ80                GarminProduct = 2806
	GarminProduct_Fr35Korea                  GarminProduct = 2814
	GarminProduct_D2charlie                  GarminProduct = 2819
	GarminProduct_VivoSmart3Apac             GarminProduct = 2831
	GarminProduct_VivoSportApac              GarminProduct = 2832
	GarminProduct_Fr935Asia                  GarminProduct = 2833
	GarminProduct_Descent                    GarminProduct = 2859
	GarminProduct_VivoFit4                   GarminProduct = 2878
	GarminProduct_Fr645                      GarminProduct = 2886
	GarminProduct_Fr645m                     GarminProduct = 2888
	GarminProduct_Fr30                       GarminProduct = 2891
	GarminProduct_Fenix5sPlus                GarminProduct = 2900
	GarminProduct_Edge130                    GarminProduct = 2909
	GarminProduct_Edge1030Asia               GarminProduct = 2924
	GarminProduct_Vivosmart4                 GarminProduct = 2927
	GarminProduct_VivoMoveHrAsia             GarminProduct = 2945
	GarminProduct_ApproachX10                GarminProduct = 2962
	GarminProduct_Fr30Asia                   GarminProduct = 2977
	GarminProduct_Vivoactive3mW              GarminProduct = 2988
	GarminProduct_Fr645Asia                  GarminProduct = 3003
	GarminProduct_Fr645mAsia                 GarminProduct = 3004
	GarminProduct_EdgeExplore                GarminProduct = 3011
	GarminProduct_Gpsmap66                   GarminProduct = 3028
	GarminProduct_ApproachS10                GarminProduct = 3049
	GarminProduct_Vivoactive3mL              GarminProduct = 3066
	GarminProduct_ApproachG80                GarminProduct = 3085
	GarminProduct_Edge130Asia                GarminProduct = 3092
	GarminProduct_Edge1030Bontrager          GarminProduct = 3095
	GarminProduct_Fenix5Plus                 GarminProduct = 3110
	GarminProduct_Fenix5xPlus                GarminProduct = 3111
	GarminProduct_Edge520Plus                GarminProduct = 3112
	GarminProduct_Fr945                      GarminProduct = 3113
	GarminProduct_Edge530                    GarminProduct = 3121
	GarminProduct_Edge830                    GarminProduct = 3122
	GarminProduct_InstinctEsports            GarminProduct = 3126
	GarminProduct_Fenix5sPlusApac            GarminProduct = 3134
	GarminProduct_Fenix5xPlusApac            GarminProduct = 3135
	GarminProduct_Edge520PlusApac            GarminProduct = 3142
	GarminProduct_DescentT1                  GarminProduct = 3143
	GarminProduct_Fr235lAsia                 GarminProduct = 3144
	GarminProduct_Fr245Asia                  GarminProduct = 3145
	GarminProduct_VivoActive3mApac           GarminProduct = 3163
	GarminProduct_Gen3Bsm                    GarminProduct = 3192
	GarminProduct_Gen3Bcm                    GarminProduct = 3193
	GarminProduct_VivoSmart4Asia             GarminProduct = 3218
	GarminProduct_Vivoactive4Small           GarminProduct = 3224
	GarminProduct_Vivoactive4Large           GarminProduct = 3225
	GarminProduct_Venu                       GarminProduct = 3226
	GarminProduct_MarqDriver                 GarminProduct = 3246
	GarminProduct_MarqAviator                GarminProduct = 3247
	GarminProduct_MarqCaptain                GarminProduct = 3248
	GarminProduct_MarqCommander              GarminProduct = 3249
	GarminProduct_MarqExpedition             GarminProduct = 3250
	GarminProduct_MarqAthlete                GarminProduct = 3251
	GarminProduct_DescentMk2                 GarminProduct = 3258
	GarminProduct_Gpsmap66i                  GarminProduct = 3284
	GarminProduct_Fenix6SSport               GarminProduct = 3287
	GarminProduct_Fenix6S                    GarminProduct = 3288
	GarminProduct_Fenix6Sport                GarminProduct = 3289
	GarminProduct_Fenix6                     GarminProduct = 3290
	GarminProduct_Fenix6x                    GarminProduct = 3291
	GarminProduct_HrmDual                    GarminProduct = 3299
	GarminProduct_HrmPro                     GarminProduct = 3300
	GarminProduct_VivoMove3Premium           GarminProduct = 3308
	GarminProduct_ApproachS40                GarminProduct = 3314
	GarminProduct_Fr245mAsia                 GarminProduct = 3321
	GarminProduct_Edge530Apac                GarminProduct = 3349
	GarminProduct_Edge830Apac                GarminProduct = 3350
	GarminProduct_VivoMove3                  GarminProduct = 3378
	GarminProduct_VivoActive4SmallAsia       GarminProduct = 3387
	GarminProduct_VivoActive4LargeAsia       GarminProduct = 3388
	GarminProduct_VivoActive4OledAsia        GarminProduct = 3389
	GarminProduct_Swim2                      GarminProduct = 3405
	GarminProduct_MarqDriverAsia             GarminProduct = 3420
	GarminProduct_MarqAviatorAsia            GarminProduct = 3421
	GarminProduct_VivoMove3Asia              GarminProduct = 3422
	GarminProduct_Fr945Asia                  GarminProduct = 3441
	GarminProduct_VivoActive3tChn            GarminProduct = 3446
	GarminProduct_MarqCaptainAsia            GarminProduct = 3448
	GarminProduct_MarqCommanderAsia          GarminProduct = 3449
	GarminProduct_MarqExpeditionAsia         GarminProduct = 3450
	GarminProduct_MarqAthleteAsia            GarminProduct = 3451
	GarminProduct_IndexSmartScale2           GarminProduct = 3461
	GarminProduct_InstinctSolar              GarminProduct = 3466
	GarminProduct_Fr45Asia                   GarminProduct = 3469
	GarminProduct_Vivoactive3Daimler         GarminProduct = 3473
	GarminProduct_LegacyRey                  GarminProduct = 3498
	GarminProduct_LegacyDarthVader           GarminProduct = 3499
	GarminProduct_LegacyCaptainMarvel        GarminProduct = 3500
	GarminProduct_LegacyFirstAvenger         GarminProduct = 3501
	GarminProduct_Fenix6sSportAsia           GarminProduct = 3512
	GarminProduct_Fenix6sAsia                GarminProduct = 3513
	GarminProduct_Fenix6SportAsia            GarminProduct = 3514
	GarminProduct_Fenix6Asia                 GarminProduct = 3515
	GarminProduct_Fenix6xAsia                GarminProduct = 3516
	GarminProduct_LegacyCaptainMarvelAsia    GarminProduct = 3535
	GarminProduct_LegacyFirstAvengerAsia     GarminProduct = 3536
	GarminProduct_LegacyReyAsia              GarminProduct = 3537
	GarminProduct_LegacyDarthVaderAsia       GarminProduct = 3538
	GarminProduct_DescentMk2s                GarminProduct = 3542
	GarminProduct_Edge130Plus                GarminProduct = 3558
	GarminProduct_Edge1030Plus               GarminProduct = 3570
	GarminProduct_Rally200                   GarminProduct = 3578
	GarminProduct_Fr745                      GarminProduct = 3589
	GarminProduct_Venusq                     GarminProduct = 3600
	GarminProduct_Lily                       GarminProduct = 3615
	GarminProduct_MarqAdventurer             GarminProduct = 3624
	GarminProduct_Enduro                     GarminProduct = 3638
	GarminProduct_Swim2Apac                  GarminProduct = 3639
	GarminProduct_MarqAdventurerAsia         GarminProduct = 3648
	GarminProduct_Fr945Lte                   GarminProduct = 3652
	GarminProduct_DescentMk2Asia             GarminProduct = 3702
	GarminProduct_Venu2                      GarminProduct = 3703
	GarminProduct_Venu2s                     GarminProduct = 3704
	GarminProduct_VenuDaimlerAsia            GarminProduct = 3737
	GarminProduct_MarqGolfer                 GarminProduct = 3739
	GarminProduct_VenuDaimler                GarminProduct = 3740
	GarminProduct_Fr745Asia                  GarminProduct = 3794
	GarminProduct_VariaRct715                GarminProduct = 3808
	GarminProduct_LilyAsia                   GarminProduct = 3809
	GarminProduct_Edge1030PlusAsia           GarminProduct = 3812
	GarminProduct_Edge130PlusAsia            GarminProduct = 3813
	GarminProduct_ApproachS12                GarminProduct = 3823
	GarminProduct_EnduroAsia                 GarminProduct = 3872
	GarminProduct_VenusqAsia                 GarminProduct = 3837
	GarminProduct_Edge1040                   GarminProduct = 3843
	GarminProduct_MarqGolferAsia             GarminProduct = 3850
	GarminProduct_Venu2Plus                  GarminProduct = 3851
	GarminProduct_Gnss                       GarminProduct = 3865
	GarminProduct_Fr55                       GarminProduct = 3869
	GarminProduct_Instinct2                  GarminProduct = 3888
	GarminProduct_Fenix7s                    GarminProduct = 3905
	GarminProduct_Fenix7                     GarminProduct = 3906
	GarminProduct_Fenix7x                    GarminProduct = 3907
	GarminProduct_Fenix7sApac                GarminProduct = 3908
	GarminProduct_Fenix7Apac                 GarminProduct = 3909
	GarminProduct_Fenix7xApac                GarminProduct = 3910
	GarminProduct_ApproachG12                GarminProduct = 3927
	GarminProduct_DescentMk2sAsia            GarminProduct = 3930
	GarminProduct_ApproachS42                GarminProduct = 3934
	GarminProduct_EpixGen2                   GarminProduct = 3943
	GarminProduct_EpixGen2Apac               GarminProduct = 3944
	GarminProduct_Venu2sAsia                 GarminProduct = 3949
	GarminProduct_Venu2Asia                  GarminProduct = 3950
	GarminProduct_Fr945LteAsia               GarminProduct = 3978
	GarminProduct_VivoMoveSport              GarminProduct = 3982
	GarminProduct_VivomoveTrend              GarminProduct = 3983
	GarminProduct_ApproachS12Asia            GarminProduct = 3986
	GarminProduct_Fr255Music                 GarminProduct = 3990
	GarminProduct_Fr255SmallMusic            GarminProduct = 3991
	GarminProduct_Fr255                      GarminProduct = 3992
	GarminProduct_Fr255Small                 GarminProduct = 3993
	GarminProduct_ApproachG12Asia            GarminProduct = 4001
	GarminProduct_ApproachS42Asia            GarminProduct = 4002
	GarminProduct_DescentG1                  GarminProduct = 4005
	GarminProduct_Venu2PlusAsia              GarminProduct = 4017
	GarminProduct_Fr955                      GarminProduct = 4024
	GarminProduct_Fr55Asia                   GarminProduct = 4033
	GarminProduct_Edge540                    GarminProduct = 4061
	GarminProduct_Edge840                    GarminProduct = 4062
	GarminProduct_Vivosmart5                 GarminProduct = 4063
	GarminProduct_Instinct2Asia              GarminProduct = 4071
	GarminProduct_MarqGen2                   GarminProduct = 4105
	GarminProduct_Venusq2                    GarminProduct = 4115
	GarminProduct_Venusq2music               GarminProduct = 4116
	GarminProduct_MarqGen2Aviator            GarminProduct = 4124
	GarminProduct_D2AirX10                   GarminProduct = 4125
	GarminProduct_HrmProPlus                 GarminProduct = 4130
	GarminProduct_DescentG1Asia              GarminProduct = 4132
	GarminProduct_Tactix7                    GarminProduct = 4135
	GarminProduct_InstinctCrossover          GarminProduct = 4155
	GarminProduct_EdgeExplore2               GarminProduct = 4169
	GarminProduct_DescentMk3                 GarminProduct = 4222
	GarminProduct_DescentMk3i                GarminProduct = 4223
	GarminProduct_ApproachS70                GarminProduct = 4233
	GarminProduct_Fr265Large                 GarminProduct = 4257
	GarminProduct_Fr265Small                 GarminProduct = 4258
	GarminProduct_Venu3                      GarminProduct = 4260
	GarminProduct_Venu3s                     GarminProduct = 4261
	GarminProduct_TacxNeoSmart               GarminProduct = 4265
	GarminProduct_TacxNeo2Smart              GarminProduct = 4266
	GarminProduct_TacxNeo2TSmart             GarminProduct = 4267
	GarminProduct_TacxNeoSmartBike           GarminProduct = 4268
	GarminProduct_TacxSatoriSmart            GarminProduct = 4269
	GarminProduct_TacxFlowSmart              GarminProduct = 4270
	GarminProduct_TacxVortexSmart            GarminProduct = 4271
	GarminProduct_TacxBushidoSmart           GarminProduct = 4272
	GarminProduct_TacxGeniusSmart            GarminProduct = 4273
	GarminProduct_TacxFluxFluxSSmart         GarminProduct = 4274
	GarminProduct_TacxFlux2Smart             GarminProduct = 4275
	GarminProduct_TacxMagnum                 GarminProduct = 4276
	GarminProduct_Edge1040Asia               GarminProduct = 4305
	GarminProduct_EpixGen2Pro42              GarminProduct = 4312
	GarminProduct_EpixGen2Pro47              GarminProduct = 4313
	GarminProduct_EpixGen2Pro51              GarminProduct = 4314
	GarminProduct_Fr965                      GarminProduct = 4315
	GarminProduct_Enduro2                    GarminProduct = 4341
	GarminProduct_Fenix7sProSolar            GarminProduct = 4374
	GarminProduct_Fenix7ProSolar             GarminProduct = 4375
	GarminProduct_Fenix7xProSolar            GarminProduct = 4376
	GarminProduct_Lily2                      GarminProduct = 4380
	GarminProduct_Instinct2x                 GarminProduct = 4394
	GarminProduct_Vivoactive5                GarminProduct = 4426
	GarminProduct_Fr165                      GarminProduct = 4432
	GarminProduct_Fr165Music                 GarminProduct = 4433
	GarminProduct_Edge1050                   GarminProduct = 4440
	GarminProduct_DescentT2                  GarminProduct = 4442
	GarminProduct_HrmFit                     GarminProduct = 4446
	GarminProduct_MarqGen2Commander          GarminProduct = 4472
	GarminProduct_LilyAthlete                GarminProduct = 4477
	GarminProduct_Fenix8Solar                GarminProduct = 4532
	GarminProduct_Fenix8SolarLarge           GarminProduct = 4533
	GarminProduct_Fenix8Small                GarminProduct = 4534
	GarminProduct_Fenix8                     GarminProduct = 4536
	GarminProduct_D2Mach1Pro                 GarminProduct = 4556
	GarminProduct_Enduro3                    GarminProduct = 4575
	GarminProduct_FenixE                     GarminProduct = 4666
	GarminProduct_Sdm4                       GarminProduct = 10007
	GarminProduct_EdgeRemote                 GarminProduct = 10014
	GarminProduct_TacxTrainingAppWin         GarminProduct = 20533
	GarminProduct_TacxTrainingAppMac         GarminProduct = 20534
	GarminProduct_TacxTrainingAppMacCatalyst GarminProduct = 20565
	GarminProduct_TrainingCenter             GarminProduct = 20119
	GarminProduct_TacxTrainingAppAndroid     GarminProduct = 30045
	GarminProduct_TacxTrainingAppIos         GarminProduct = 30046
	GarminProduct_TacxTrainingAppLegacy      GarminProduct = 30047
	GarminProduct_ConnectiqSimulator         GarminProduct = 65531
	GarminProduct_AndroidAntplusPlugin       GarminProduct = 65532
	GarminProduct_Connect                    GarminProduct = 65534
	GarminProduct_Invalid                    GarminProduct = 65535
)

var GarminProductNames = map[GarminProduct]string{
	GarminProduct_Hrm1:                       "Hrm1",
	GarminProduct_Axh01:                      "Axh01",
	GarminProduct_Axb01:                      "Axb01",
	GarminProduct_Axb02:                      "Axb02",
	GarminProduct_Hrm2ss:                     "Hrm2ss",
	GarminProduct_DsiAlf02:                   "DsiAlf02",
	GarminProduct_Hrm3ss:                     "Hrm3ss",
	GarminProduct_HrmRunSingleByteProductId:  "HrmRunSingleByteProductId",
	GarminProduct_Bsm:                        "Bsm",
	GarminProduct_Bcm:                        "Bcm",
	GarminProduct_Axs01:                      "Axs01",
	GarminProduct_HrmTriSingleByteProductId:  "HrmTriSingleByteProductId",
	GarminProduct_Hrm4RunSingleByteProductId: "Hrm4RunSingleByteProductId",
	GarminProduct_Fr225SingleByteProductId:   "Fr225SingleByteProductId",
	GarminProduct_Gen3BsmSingleByteProductId: "Gen3BsmSingleByteProductId",
	GarminProduct_Gen3BcmSingleByteProductId: "Gen3BcmSingleByteProductId",
	GarminProduct_HrmFitSingleByteProductId:  "HrmFitSingleByteProductId",
	GarminProduct_OHR:                        "OHR",
	GarminProduct_Fr301China:                 "Fr301China",
	GarminProduct_Fr301Japan:                 "Fr301Japan",
	GarminProduct_Fr301Korea:                 "Fr301Korea",
	GarminProduct_Fr301Taiwan:                "Fr301Taiwan",
	GarminProduct_Fr405:                      "Fr405",
	GarminProduct_Fr50:                       "Fr50",
	GarminProduct_Fr405Japan:                 "Fr405Japan",
	GarminProduct_Fr60:                       "Fr60",
	GarminProduct_DsiAlf01:                   "DsiAlf01",
	GarminProduct_Fr310xt:                    "Fr310xt",
	GarminProduct_Edge500:                    "Edge500",
	GarminProduct_Fr110:                      "Fr110",
	GarminProduct_Edge800:                    "Edge800",
	GarminProduct_Edge500Taiwan:              "Edge500Taiwan",
	GarminProduct_Edge500Japan:               "Edge500Japan",
	GarminProduct_Chirp:                      "Chirp",
	GarminProduct_Fr110Japan:                 "Fr110Japan",
	GarminProduct_Edge200:                    "Edge200",
	GarminProduct_Fr910xt:                    "Fr910xt",
	GarminProduct_Edge800Taiwan:              "Edge800Taiwan",
	GarminProduct_Edge800Japan:               "Edge800Japan",
	GarminProduct_Alf04:                      "Alf04",
	GarminProduct_Fr610:                      "Fr610",
	GarminProduct_Fr210Japan:                 "Fr210Japan",
	GarminProduct_VectorSs:                   "VectorSs",
	GarminProduct_VectorCp:                   "VectorCp",
	GarminProduct_Edge800China:               "Edge800China",
	GarminProduct_Edge500China:               "Edge500China",
	GarminProduct_ApproachG10:                "ApproachG10",
	GarminProduct_Fr610Japan:                 "Fr610Japan",
	GarminProduct_Edge500Korea:               "Edge500Korea",
	GarminProduct_Fr70:                       "Fr70",
	GarminProduct_Fr310xt4t:                  "Fr310xt4t",
	GarminProduct_Amx:                        "Amx",
	GarminProduct_Fr10:                       "Fr10",
	GarminProduct_Edge800Korea:               "Edge800Korea",
	GarminProduct_Swim:                       "Swim",
	GarminProduct_Fr910xtChina:               "Fr910xtChina",
	GarminProduct_Fenix:                      "Fenix",
	GarminProduct_Edge200Taiwan:              "Edge200Taiwan",
	GarminProduct_Edge510:                    "Edge510",
	GarminProduct_Edge810:                    "Edge810",
	GarminProduct_Tempe:                      "Tempe",
	GarminProduct_Fr910xtJapan:               "Fr910xtJapan",
	GarminProduct_Fr620:                      "Fr620",
	GarminProduct_Fr220:                      "Fr220",
	GarminProduct_Fr910xtKorea:               "Fr910xtKorea",
	GarminProduct_Fr10Japan:                  "Fr10Japan",
	GarminProduct_Edge810Japan:               "Edge810Japan",
	GarminProduct_VirbElite:                  "VirbElite",
	GarminProduct_EdgeTouring:                "EdgeTouring",
	GarminProduct_Edge510Japan:               "Edge510Japan",
	GarminProduct_HrmTri:                     "HrmTri",
	GarminProduct_HrmRun:                     "HrmRun",
	GarminProduct_Fr920xt:                    "Fr920xt",
	GarminProduct_Edge510Asia:                "Edge510Asia",
	GarminProduct_Edge810China:               "Edge810China",
	GarminProduct_Edge810Taiwan:              "Edge810Taiwan",
	GarminProduct_Edge1000:                   "Edge1000",
	GarminProduct_VivoFit:                    "VivoFit",
	GarminProduct_VirbRemote:                 "VirbRemote",
	GarminProduct_VivoKi:                     "VivoKi",
	GarminProduct_Fr15:                       "Fr15",
	GarminProduct_VivoActive:                 "VivoActive",
	GarminProduct_Edge510Korea:               "Edge510Korea",
	GarminProduct_Fr620Japan:                 "Fr620Japan",
	GarminProduct_Fr620China:                 "Fr620China",
	GarminProduct_Fr220Japan:                 "Fr220Japan",
	GarminProduct_Fr220China:                 "Fr220China",
	GarminProduct_ApproachS6:                 "ApproachS6",
	GarminProduct_VivoSmart:                  "VivoSmart",
	GarminProduct_Fenix2:                     "Fenix2",
	GarminProduct_Epix:                       "Epix",
	GarminProduct_Fenix3:                     "Fenix3",
	GarminProduct_Edge1000Taiwan:             "Edge1000Taiwan",
	GarminProduct_Edge1000Japan:              "Edge1000Japan",
	GarminProduct_Fr15Japan:                  "Fr15Japan",
	GarminProduct_Edge520:                    "Edge520",
	GarminProduct_Edge1000China:              "Edge1000China",
	GarminProduct_Fr620Russia:                "Fr620Russia",
	GarminProduct_Fr220Russia:                "Fr220Russia",
	GarminProduct_VectorS:                    "VectorS",
	GarminProduct_Edge1000Korea:              "Edge1000Korea",
	GarminProduct_Fr920xtTaiwan:              "Fr920xtTaiwan",
	GarminProduct_Fr920xtChina:               "Fr920xtChina",
	GarminProduct_Fr920xtJapan:               "Fr920xtJapan",
	GarminProduct_Virbx:                      "Virbx",
	GarminProduct_VivoSmartApac:              "VivoSmartApac",
	GarminProduct_EtrexTouch:                 "EtrexTouch",
	GarminProduct_Edge25:                     "Edge25",
	GarminProduct_Fr25:                       "Fr25",
	GarminProduct_VivoFit2:                   "VivoFit2",
	GarminProduct_Fr225:                      "Fr225",
	GarminProduct_Fr630:                      "Fr630",
	GarminProduct_Fr230:                      "Fr230",
	GarminProduct_Fr735xt:                    "Fr735xt",
	GarminProduct_VivoActiveApac:             "VivoActiveApac",
	GarminProduct_Vector2:                    "Vector2",
	GarminProduct_Vector2s:                   "Vector2s",
	GarminProduct_Virbxe:                     "Virbxe",
	GarminProduct_Fr620Taiwan:                "Fr620Taiwan",
	GarminProduct_Fr220Taiwan:                "Fr220Taiwan",
	GarminProduct_Truswing:                   "Truswing",
	GarminProduct_D2airvenu:                  "D2airvenu",
	GarminProduct_Fenix3China:                "Fenix3China",
	GarminProduct_Fenix3Twn:                  "Fenix3Twn",
	GarminProduct_VariaHeadlight:             "VariaHeadlight",
	GarminProduct_VariaTaillightOld:          "VariaTaillightOld",
	GarminProduct_EdgeExplore1000:            "EdgeExplore1000",
	GarminProduct_Fr225Asia:                  "Fr225Asia",
	GarminProduct_VariaRadarTaillight:        "VariaRadarTaillight",
	GarminProduct_VariaRadarDisplay:          "VariaRadarDisplay",
	GarminProduct_Edge20:                     "Edge20",
	GarminProduct_Edge520Asia:                "Edge520Asia",
	GarminProduct_Edge520Japan:               "Edge520Japan",
	GarminProduct_D2Bravo:                    "D2Bravo",
	GarminProduct_ApproachS20:                "ApproachS20",
	GarminProduct_VivoSmart2:                 "VivoSmart2",
	GarminProduct_Edge1000Thai:               "Edge1000Thai",
	GarminProduct_VariaRemote:                "VariaRemote",
	GarminProduct_Edge25Asia:                 "Edge25Asia",
	GarminProduct_Edge25Jpn:                  "Edge25Jpn",
	GarminProduct_Edge20Asia:                 "Edge20Asia",
	GarminProduct_ApproachX40:                "ApproachX40",
	GarminProduct_Fenix3Japan:                "Fenix3Japan",
	GarminProduct_VivoSmartEmea:              "VivoSmartEmea",
	GarminProduct_Fr630Asia:                  "Fr630Asia",
	GarminProduct_Fr630Jpn:                   "Fr630Jpn",
	GarminProduct_Fr230Jpn:                   "Fr230Jpn",
	GarminProduct_Hrm4Run:                    "Hrm4Run",
	GarminProduct_EpixJapan:                  "EpixJapan",
	GarminProduct_VivoActiveHr:               "VivoActiveHr",
	GarminProduct_VivoSmartGpsHr:             "VivoSmartGpsHr",
	GarminProduct_VivoSmartHr:                "VivoSmartHr",
	GarminProduct_VivoSmartHrAsia:            "VivoSmartHrAsia",
	GarminProduct_VivoSmartGpsHrAsia:         "VivoSmartGpsHrAsia",
	GarminProduct_VivoMove:                   "VivoMove",
	GarminProduct_VariaTaillight:             "VariaTaillight",
	GarminProduct_Fr235Asia:                  "Fr235Asia",
	GarminProduct_Fr235Japan:                 "Fr235Japan",
	GarminProduct_VariaVision:                "VariaVision",
	GarminProduct_VivoFit3:                   "VivoFit3",
	GarminProduct_Fenix3Korea:                "Fenix3Korea",
	GarminProduct_Fenix3Sea:                  "Fenix3Sea",
	GarminProduct_Fenix3Hr:                   "Fenix3Hr",
	GarminProduct_VirbUltra30:                "VirbUltra30",
	GarminProduct_IndexSmartScale:            "IndexSmartScale",
	GarminProduct_Fr235:                      "Fr235",
	GarminProduct_Fenix3Chronos:              "Fenix3Chronos",
	GarminProduct_Oregon7xx:                  "Oregon7xx",
	GarminProduct_Rino7xx:                    "Rino7xx",
	GarminProduct_EpixKorea:                  "EpixKorea",
	GarminProduct_Fenix3HrChn:                "Fenix3HrChn",
	GarminProduct_Fenix3HrTwn:                "Fenix3HrTwn",
	GarminProduct_Fenix3HrJpn:                "Fenix3HrJpn",
	GarminProduct_Fenix3HrSea:                "Fenix3HrSea",
	GarminProduct_Fenix3HrKor:                "Fenix3HrKor",
	GarminProduct_Nautix:                     "Nautix",
	GarminProduct_VivoActiveHrApac:           "VivoActiveHrApac",
	GarminProduct_Fr35:                       "Fr35",
	GarminProduct_Oregon7xxWw:                "Oregon7xxWw",
	GarminProduct_Edge820:                    "Edge820",
	GarminProduct_EdgeExplore820:             "EdgeExplore820",
	GarminProduct_Fr735xtApac:                "Fr735xtApac",
	GarminProduct_Fr735xtJapan:               "Fr735xtJapan",
	GarminProduct_Fenix5s:                    "Fenix5s",
	GarminProduct_D2BravoTitanium:            "D2BravoTitanium",
	GarminProduct_VariaUt800:                 "VariaUt800",
	GarminProduct_RunningDynamicsPod:         "RunningDynamicsPod",
	GarminProduct_Edge820China:               "Edge820China",
	GarminProduct_Edge820Japan:               "Edge820Japan",
	GarminProduct_Fenix5x:                    "Fenix5x",
	GarminProduct_VivoFitJr:                  "VivoFitJr",
	GarminProduct_VivoSmart3:                 "VivoSmart3",
	GarminProduct_VivoSport:                  "VivoSport",
	GarminProduct_Edge820Taiwan:              "Edge820Taiwan",
	GarminProduct_Edge820Korea:               "Edge820Korea",
	GarminProduct_Edge820Sea:                 "Edge820Sea",
	GarminProduct_Fr35Hebrew:                 "Fr35Hebrew",
	GarminProduct_ApproachS60:                "ApproachS60",
	GarminProduct_Fr35Apac:                   "Fr35Apac",
	GarminProduct_Fr35Japan:                  "Fr35Japan",
	GarminProduct_Fenix3ChronosAsia:          "Fenix3ChronosAsia",
	GarminProduct_Virb360:                    "Virb360",
	GarminProduct_Fr935:                      "Fr935",
	GarminProduct_Fenix5:                     "Fenix5",
	GarminProduct_Vivoactive3:                "Vivoactive3",
	GarminProduct_Fr235ChinaNfc:              "Fr235ChinaNfc",
	GarminProduct_Foretrex601701:             "Foretrex601701",
	GarminProduct_VivoMoveHr:                 "VivoMoveHr",
	GarminProduct_Edge1030:                   "Edge1030",
	GarminProduct_Fr35Sea:                    "Fr35Sea",
	GarminProduct_Vector3:                    "Vector3",
	GarminProduct_Fenix5Asia:                 "Fenix5Asia",
	GarminProduct_Fenix5sAsia:                "Fenix5sAsia",
	GarminProduct_Fenix5xAsia:                "Fenix5xAsia",
	GarminProduct_ApproachZ80:                "ApproachZ80",
	GarminProduct_Fr35Korea:                  "Fr35Korea",
	GarminProduct_D2charlie:                  "D2charlie",
	GarminProduct_VivoSmart3Apac:             "VivoSmart3Apac",
	GarminProduct_VivoSportApac:              "VivoSportApac",
	GarminProduct_Fr935Asia:                  "Fr935Asia",
	GarminProduct_Descent:                    "Descent",
	GarminProduct_VivoFit4:                   "VivoFit4",
	GarminProduct_Fr645:                      "Fr645",
	GarminProduct_Fr645m:                     "Fr645m",
	GarminProduct_Fr30:                       "Fr30",
	GarminProduct_Fenix5sPlus:                "Fenix5sPlus",
	GarminProduct_Edge130:                    "Edge130",
	GarminProduct_Edge1030Asia:               "Edge1030Asia",
	GarminProduct_Vivosmart4:                 "Vivosmart4",
	GarminProduct_VivoMoveHrAsia:             "VivoMoveHrAsia",
	GarminProduct_ApproachX10:                "ApproachX10",
	GarminProduct_Fr30Asia:                   "Fr30Asia",
	GarminProduct_Vivoactive3mW:              "Vivoactive3mW",
	GarminProduct_Fr645Asia:                  "Fr645Asia",
	GarminProduct_Fr645mAsia:                 "Fr645mAsia",
	GarminProduct_EdgeExplore:                "EdgeExplore",
	GarminProduct_Gpsmap66:                   "Gpsmap66",
	GarminProduct_ApproachS10:                "ApproachS10",
	GarminProduct_Vivoactive3mL:              "Vivoactive3mL",
	GarminProduct_ApproachG80:                "ApproachG80",
	GarminProduct_Edge130Asia:                "Edge130Asia",
	GarminProduct_Edge1030Bontrager:          "Edge1030Bontrager",
	GarminProduct_Fenix5Plus:                 "Fenix5Plus",
	GarminProduct_Fenix5xPlus:                "Fenix5xPlus",
	GarminProduct_Edge520Plus:                "Edge520Plus",
	GarminProduct_Fr945:                      "Fr945",
	GarminProduct_Edge530:                    "Edge530",
	GarminProduct_Edge830:                    "Edge830",
	GarminProduct_InstinctEsports:            "InstinctEsports",
	GarminProduct_Fenix5sPlusApac:            "Fenix5sPlusApac",
	GarminProduct_Fenix5xPlusApac:            "Fenix5xPlusApac",
	GarminProduct_Edge520PlusApac:            "Edge520PlusApac",
	GarminProduct_DescentT1:                  "DescentT1",
	GarminProduct_Fr235lAsia:                 "Fr235lAsia",
	GarminProduct_Fr245Asia:                  "Fr245Asia",
	GarminProduct_VivoActive3mApac:           "VivoActive3mApac",
	GarminProduct_Gen3Bsm:                    "Gen3Bsm",
	GarminProduct_Gen3Bcm:                    "Gen3Bcm",
	GarminProduct_VivoSmart4Asia:             "VivoSmart4Asia",
	GarminProduct_Vivoactive4Small:           "Vivoactive4Small",
	GarminProduct_Vivoactive4Large:           "Vivoactive4Large",
	GarminProduct_Venu:                       "Venu",
	GarminProduct_MarqDriver:                 "MarqDriver",
	GarminProduct_MarqAviator:                "MarqAviator",
	GarminProduct_MarqCaptain:                "MarqCaptain",
	GarminProduct_MarqCommander:              "MarqCommander",
	GarminProduct_MarqExpedition:             "MarqExpedition",
	GarminProduct_MarqAthlete:                "MarqAthlete",
	GarminProduct_DescentMk2:                 "DescentMk2",
	GarminProduct_Gpsmap66i:                  "Gpsmap66i",
	GarminProduct_Fenix6SSport:               "Fenix6SSport",
	GarminProduct_Fenix6S:                    "Fenix6S",
	GarminProduct_Fenix6Sport:                "Fenix6Sport",
	GarminProduct_Fenix6:                     "Fenix6",
	GarminProduct_Fenix6x:                    "Fenix6x",
	GarminProduct_HrmDual:                    "HrmDual",
	GarminProduct_HrmPro:                     "HrmPro",
	GarminProduct_VivoMove3Premium:           "VivoMove3Premium",
	GarminProduct_ApproachS40:                "ApproachS40",
	GarminProduct_Fr245mAsia:                 "Fr245mAsia",
	GarminProduct_Edge530Apac:                "Edge530Apac",
	GarminProduct_Edge830Apac:                "Edge830Apac",
	GarminProduct_VivoMove3:                  "VivoMove3",
	GarminProduct_VivoActive4SmallAsia:       "VivoActive4SmallAsia",
	GarminProduct_VivoActive4LargeAsia:       "VivoActive4LargeAsia",
	GarminProduct_VivoActive4OledAsia:        "VivoActive4OledAsia",
	GarminProduct_Swim2:                      "Swim2",
	GarminProduct_MarqDriverAsia:             "MarqDriverAsia",
	GarminProduct_MarqAviatorAsia:            "MarqAviatorAsia",
	GarminProduct_VivoMove3Asia:              "VivoMove3Asia",
	GarminProduct_Fr945Asia:                  "Fr945Asia",
	GarminProduct_VivoActive3tChn:            "VivoActive3tChn",
	GarminProduct_MarqCaptainAsia:            "MarqCaptainAsia",
	GarminProduct_MarqCommanderAsia:          "MarqCommanderAsia",
	GarminProduct_MarqExpeditionAsia:         "MarqExpeditionAsia",
	GarminProduct_MarqAthleteAsia:            "MarqAthleteAsia",
	GarminProduct_IndexSmartScale2:           "IndexSmartScale2",
	GarminProduct_InstinctSolar:              "InstinctSolar",
	GarminProduct_Fr45Asia:                   "Fr45Asia",
	GarminProduct_Vivoactive3Daimler:         "Vivoactive3Daimler",
	GarminProduct_LegacyRey:                  "LegacyRey",
	GarminProduct_LegacyDarthVader:           "LegacyDarthVader",
	GarminProduct_LegacyCaptainMarvel:        "LegacyCaptainMarvel",
	GarminProduct_LegacyFirstAvenger:         "LegacyFirstAvenger",
	GarminProduct_Fenix6sSportAsia:           "Fenix6sSportAsia",
	GarminProduct_Fenix6sAsia:                "Fenix6sAsia",
	GarminProduct_Fenix6SportAsia:            "Fenix6SportAsia",
	GarminProduct_Fenix6Asia:                 "Fenix6Asia",
	GarminProduct_Fenix6xAsia:                "Fenix6xAsia",
	GarminProduct_LegacyCaptainMarvelAsia:    "LegacyCaptainMarvelAsia",
	GarminProduct_LegacyFirstAvengerAsia:     "LegacyFirstAvengerAsia",
	GarminProduct_LegacyReyAsia:              "LegacyReyAsia",
	GarminProduct_LegacyDarthVaderAsia:       "LegacyDarthVaderAsia",
	GarminProduct_DescentMk2s:                "DescentMk2s",
	GarminProduct_Edge130Plus:                "Edge130Plus",
	GarminProduct_Edge1030Plus:               "Edge1030Plus",
	GarminProduct_Rally200:                   "Rally200",
	GarminProduct_Fr745:                      "Fr745",
	GarminProduct_Venusq:                     "Venusq",
	GarminProduct_Lily:                       "Lily",
	GarminProduct_MarqAdventurer:             "MarqAdventurer",
	GarminProduct_Enduro:                     "Enduro",
	GarminProduct_Swim2Apac:                  "Swim2Apac",
	GarminProduct_MarqAdventurerAsia:         "MarqAdventurerAsia",
	GarminProduct_Fr945Lte:                   "Fr945Lte",
	GarminProduct_DescentMk2Asia:             "DescentMk2Asia",
	GarminProduct_Venu2:                      "Venu2",
	GarminProduct_Venu2s:                     "Venu2s",
	GarminProduct_VenuDaimlerAsia:            "VenuDaimlerAsia",
	GarminProduct_MarqGolfer:                 "MarqGolfer",
	GarminProduct_VenuDaimler:                "VenuDaimler",
	GarminProduct_Fr745Asia:                  "Fr745Asia",
	GarminProduct_VariaRct715:                "VariaRct715",
	GarminProduct_LilyAsia:                   "LilyAsia",
	GarminProduct_Edge1030PlusAsia:           "Edge1030PlusAsia",
	GarminProduct_Edge130PlusAsia:            "Edge130PlusAsia",
	GarminProduct_ApproachS12:                "ApproachS12",
	GarminProduct_EnduroAsia:                 "EnduroAsia",
	GarminProduct_VenusqAsia:                 "VenusqAsia",
	GarminProduct_Edge1040:                   "Edge1040",
	GarminProduct_MarqGolferAsia:             "MarqGolferAsia",
	GarminProduct_Venu2Plus:                  "Venu2Plus",
	GarminProduct_Gnss:                       "Gnss",
	GarminProduct_Fr55:                       "Fr55",
	GarminProduct_Instinct2:                  "Instinct2",
	GarminProduct_Fenix7s:                    "Fenix7s",
	GarminProduct_Fenix7:                     "Fenix7",
	GarminProduct_Fenix7x:                    "Fenix7x",
	GarminProduct_Fenix7sApac:                "Fenix7sApac",
	GarminProduct_Fenix7Apac:                 "Fenix7Apac",
	GarminProduct_Fenix7xApac:                "Fenix7xApac",
	GarminProduct_ApproachG12:                "ApproachG12",
	GarminProduct_DescentMk2sAsia:            "DescentMk2sAsia",
	GarminProduct_ApproachS42:                "ApproachS42",
	GarminProduct_EpixGen2:                   "EpixGen2",
	GarminProduct_EpixGen2Apac:               "EpixGen2Apac",
	GarminProduct_Venu2sAsia:                 "Venu2sAsia",
	GarminProduct_Venu2Asia:                  "Venu2Asia",
	GarminProduct_Fr945LteAsia:               "Fr945LteAsia",
	GarminProduct_VivoMoveSport:              "VivoMoveSport",
	GarminProduct_VivomoveTrend:              "VivomoveTrend",
	GarminProduct_ApproachS12Asia:            "ApproachS12Asia",
	GarminProduct_Fr255Music:                 "Fr255Music",
	GarminProduct_Fr255SmallMusic:            "Fr255SmallMusic",
	GarminProduct_Fr255:                      "Fr255",
	GarminProduct_Fr255Small:                 "Fr255Small",
	GarminProduct_ApproachG12Asia:            "ApproachG12Asia",
	GarminProduct_ApproachS42Asia:            "ApproachS42Asia",
	GarminProduct_DescentG1:                  "DescentG1",
	GarminProduct_Venu2PlusAsia:              "Venu2PlusAsia",
	GarminProduct_Fr955:                      "Fr955",
	GarminProduct_Fr55Asia:                   "Fr55Asia",
	GarminProduct_Edge540:                    "Edge540",
	GarminProduct_Edge840:                    "Edge840",
	GarminProduct_Vivosmart5:                 "Vivosmart5",
	GarminProduct_Instinct2Asia:              "Instinct2Asia",
	GarminProduct_MarqGen2:                   "MarqGen2",
	GarminProduct_Venusq2:                    "Venusq2",
	GarminProduct_Venusq2music:               "Venusq2music",
	GarminProduct_MarqGen2Aviator:            "MarqGen2Aviator",
	GarminProduct_D2AirX10:                   "D2AirX10",
	GarminProduct_HrmProPlus:                 "HrmProPlus",
	GarminProduct_DescentG1Asia:              "DescentG1Asia",
	GarminProduct_Tactix7:                    "Tactix7",
	GarminProduct_InstinctCrossover:          "InstinctCrossover",
	GarminProduct_EdgeExplore2:               "EdgeExplore2",
	GarminProduct_DescentMk3:                 "DescentMk3",
	GarminProduct_DescentMk3i:                "DescentMk3i",
	GarminProduct_ApproachS70:                "ApproachS70",
	GarminProduct_Fr265Large:                 "Fr265Large",
	GarminProduct_Fr265Small:                 "Fr265Small",
	GarminProduct_Venu3:                      "Venu3",
	GarminProduct_Venu3s:                     "Venu3s",
	GarminProduct_TacxNeoSmart:               "TacxNeoSmart",
	GarminProduct_TacxNeo2Smart:              "TacxNeo2Smart",
	GarminProduct_TacxNeo2TSmart:             "TacxNeo2TSmart",
	GarminProduct_TacxNeoSmartBike:           "TacxNeoSmartBike",
	GarminProduct_TacxSatoriSmart:            "TacxSatoriSmart",
	GarminProduct_TacxFlowSmart:              "TacxFlowSmart",
	GarminProduct_TacxVortexSmart:            "TacxVortexSmart",
	GarminProduct_TacxBushidoSmart:           "TacxBushidoSmart",
	GarminProduct_TacxGeniusSmart:            "TacxGeniusSmart",
	GarminProduct_TacxFluxFluxSSmart:         "TacxFluxFluxSSmart",
	GarminProduct_TacxFlux2Smart:             "TacxFlux2Smart",
	GarminProduct_TacxMagnum:                 "TacxMagnum",
	GarminProduct_Edge1040Asia:               "Edge1040Asia",
	GarminProduct_EpixGen2Pro42:              "EpixGen2Pro42",
	GarminProduct_EpixGen2Pro47:              "EpixGen2Pro47",
	GarminProduct_EpixGen2Pro51:              "EpixGen2Pro51",
	GarminProduct_Fr965:                      "Fr965",
	GarminProduct_Enduro2:                    "Enduro2",
	GarminProduct_Fenix7sProSolar:            "Fenix7sProSolar",
	GarminProduct_Fenix7ProSolar:             "Fenix7ProSolar",
	GarminProduct_Fenix7xProSolar:            "Fenix7xProSolar",
	GarminProduct_Lily2:                      "Lily2",
	GarminProduct_Instinct2x:                 "Instinct2x",
	GarminProduct_Vivoactive5:                "Vivoactive5",
	GarminProduct_Fr165:                      "Fr165",
	GarminProduct_Fr165Music:                 "Fr165Music",
	GarminProduct_Edge1050:                   "Edge1050",
	GarminProduct_DescentT2:                  "DescentT2",
	GarminProduct_HrmFit:                     "HrmFit",
	GarminProduct_MarqGen2Commander:          "MarqGen2Commander",
	GarminProduct_LilyAthlete:                "LilyAthlete",
	GarminProduct_Fenix8Solar:                "Fenix8Solar",
	GarminProduct_Fenix8SolarLarge:           "Fenix8SolarLarge",
	GarminProduct_Fenix8Small:                "Fenix8Small",
	GarminProduct_Fenix8:                     "Fenix8",
	GarminProduct_D2Mach1Pro:                 "D2Mach1Pro",
	GarminProduct_Enduro3:                    "Enduro3",
	GarminProduct_FenixE:                     "FenixE",
	GarminProduct_Sdm4:                       "Sdm4",
	GarminProduct_EdgeRemote:                 "EdgeRemote",
	GarminProduct_TacxTrainingAppWin:         "TacxTrainingAppWin",
	GarminProduct_TacxTrainingAppMac:         "TacxTrainingAppMac",
	GarminProduct_TacxTrainingAppMacCatalyst: "TacxTrainingAppMacCatalyst",
	GarminProduct_TrainingCenter:             "TrainingCenter",
	GarminProduct_TacxTrainingAppAndroid:     "TacxTrainingAppAndroid",
	GarminProduct_TacxTrainingAppIos:         "TacxTrainingAppIos",
	GarminProduct_TacxTrainingAppLegacy:      "TacxTrainingAppLegacy",
	GarminProduct_ConnectiqSimulator:         "ConnectiqSimulator",
	GarminProduct_AndroidAntplusPlugin:       "AndroidAntplusPlugin",
	GarminProduct_Connect:                    "Connect",
}

func (k GarminProduct) String() string {
	if uint(k) < uint(len(GarminProductNames)) {
		if v, ok := GarminProductNames[k]; ok {
			return v
		}
	}
	return "GarminProduct(" + strconv.Itoa(int(k)) + ")"
}

var GarminProductValues = map[string]GarminProduct{
	"Hrm1":                       GarminProduct_Hrm1,
	"Axh01":                      GarminProduct_Axh01,
	"Axb01":                      GarminProduct_Axb01,
	"Axb02":                      GarminProduct_Axb02,
	"Hrm2ss":                     GarminProduct_Hrm2ss,
	"DsiAlf02":                   GarminProduct_DsiAlf02,
	"Hrm3ss":                     GarminProduct_Hrm3ss,
	"HrmRunSingleByteProductId":  GarminProduct_HrmRunSingleByteProductId,
	"Bsm":                        GarminProduct_Bsm,
	"Bcm":                        GarminProduct_Bcm,
	"Axs01":                      GarminProduct_Axs01,
	"HrmTriSingleByteProductId":  GarminProduct_HrmTriSingleByteProductId,
	"Hrm4RunSingleByteProductId": GarminProduct_Hrm4RunSingleByteProductId,
	"Fr225SingleByteProductId":   GarminProduct_Fr225SingleByteProductId,
	"Gen3BsmSingleByteProductId": GarminProduct_Gen3BsmSingleByteProductId,
	"Gen3BcmSingleByteProductId": GarminProduct_Gen3BcmSingleByteProductId,
	"HrmFitSingleByteProductId":  GarminProduct_HrmFitSingleByteProductId,
	"OHR":                        GarminProduct_OHR,
	"Fr301China":                 GarminProduct_Fr301China,
	"Fr301Japan":                 GarminProduct_Fr301Japan,
	"Fr301Korea":                 GarminProduct_Fr301Korea,
	"Fr301Taiwan":                GarminProduct_Fr301Taiwan,
	"Fr405":                      GarminProduct_Fr405,
	"Fr50":                       GarminProduct_Fr50,
	"Fr405Japan":                 GarminProduct_Fr405Japan,
	"Fr60":                       GarminProduct_Fr60,
	"DsiAlf01":                   GarminProduct_DsiAlf01,
	"Fr310xt":                    GarminProduct_Fr310xt,
	"Edge500":                    GarminProduct_Edge500,
	"Fr110":                      GarminProduct_Fr110,
	"Edge800":                    GarminProduct_Edge800,
	"Edge500Taiwan":              GarminProduct_Edge500Taiwan,
	"Edge500Japan":               GarminProduct_Edge500Japan,
	"Chirp":                      GarminProduct_Chirp,
	"Fr110Japan":                 GarminProduct_Fr110Japan,
	"Edge200":                    GarminProduct_Edge200,
	"Fr910xt":                    GarminProduct_Fr910xt,
	"Edge800Taiwan":              GarminProduct_Edge800Taiwan,
	"Edge800Japan":               GarminProduct_Edge800Japan,
	"Alf04":                      GarminProduct_Alf04,
	"Fr610":                      GarminProduct_Fr610,
	"Fr210Japan":                 GarminProduct_Fr210Japan,
	"VectorSs":                   GarminProduct_VectorSs,
	"VectorCp":                   GarminProduct_VectorCp,
	"Edge800China":               GarminProduct_Edge800China,
	"Edge500China":               GarminProduct_Edge500China,
	"ApproachG10":                GarminProduct_ApproachG10,
	"Fr610Japan":                 GarminProduct_Fr610Japan,
	"Edge500Korea":               GarminProduct_Edge500Korea,
	"Fr70":                       GarminProduct_Fr70,
	"Fr310xt4t":                  GarminProduct_Fr310xt4t,
	"Amx":                        GarminProduct_Amx,
	"Fr10":                       GarminProduct_Fr10,
	"Edge800Korea":               GarminProduct_Edge800Korea,
	"Swim":                       GarminProduct_Swim,
	"Fr910xtChina":               GarminProduct_Fr910xtChina,
	"Fenix":                      GarminProduct_Fenix,
	"Edge200Taiwan":              GarminProduct_Edge200Taiwan,
	"Edge510":                    GarminProduct_Edge510,
	"Edge810":                    GarminProduct_Edge810,
	"Tempe":                      GarminProduct_Tempe,
	"Fr910xtJapan":               GarminProduct_Fr910xtJapan,
	"Fr620":                      GarminProduct_Fr620,
	"Fr220":                      GarminProduct_Fr220,
	"Fr910xtKorea":               GarminProduct_Fr910xtKorea,
	"Fr10Japan":                  GarminProduct_Fr10Japan,
	"Edge810Japan":               GarminProduct_Edge810Japan,
	"VirbElite":                  GarminProduct_VirbElite,
	"EdgeTouring":                GarminProduct_EdgeTouring,
	"Edge510Japan":               GarminProduct_Edge510Japan,
	"HrmTri":                     GarminProduct_HrmTri,
	"HrmRun":                     GarminProduct_HrmRun,
	"Fr920xt":                    GarminProduct_Fr920xt,
	"Edge510Asia":                GarminProduct_Edge510Asia,
	"Edge810China":               GarminProduct_Edge810China,
	"Edge810Taiwan":              GarminProduct_Edge810Taiwan,
	"Edge1000":                   GarminProduct_Edge1000,
	"VivoFit":                    GarminProduct_VivoFit,
	"VirbRemote":                 GarminProduct_VirbRemote,
	"VivoKi":                     GarminProduct_VivoKi,
	"Fr15":                       GarminProduct_Fr15,
	"VivoActive":                 GarminProduct_VivoActive,
	"Edge510Korea":               GarminProduct_Edge510Korea,
	"Fr620Japan":                 GarminProduct_Fr620Japan,
	"Fr620China":                 GarminProduct_Fr620China,
	"Fr220Japan":                 GarminProduct_Fr220Japan,
	"Fr220China":                 GarminProduct_Fr220China,
	"ApproachS6":                 GarminProduct_ApproachS6,
	"VivoSmart":                  GarminProduct_VivoSmart,
	"Fenix2":                     GarminProduct_Fenix2,
	"Epix":                       GarminProduct_Epix,
	"Fenix3":                     GarminProduct_Fenix3,
	"Edge1000Taiwan":             GarminProduct_Edge1000Taiwan,
	"Edge1000Japan":              GarminProduct_Edge1000Japan,
	"Fr15Japan":                  GarminProduct_Fr15Japan,
	"Edge520":                    GarminProduct_Edge520,
	"Edge1000China":              GarminProduct_Edge1000China,
	"Fr620Russia":                GarminProduct_Fr620Russia,
	"Fr220Russia":                GarminProduct_Fr220Russia,
	"VectorS":                    GarminProduct_VectorS,
	"Edge1000Korea":              GarminProduct_Edge1000Korea,
	"Fr920xtTaiwan":              GarminProduct_Fr920xtTaiwan,
	"Fr920xtChina":               GarminProduct_Fr920xtChina,
	"Fr920xtJapan":               GarminProduct_Fr920xtJapan,
	"Virbx":                      GarminProduct_Virbx,
	"VivoSmartApac":              GarminProduct_VivoSmartApac,
	"EtrexTouch":                 GarminProduct_EtrexTouch,
	"Edge25":                     GarminProduct_Edge25,
	"Fr25":                       GarminProduct_Fr25,
	"VivoFit2":                   GarminProduct_VivoFit2,
	"Fr225":                      GarminProduct_Fr225,
	"Fr630":                      GarminProduct_Fr630,
	"Fr230":                      GarminProduct_Fr230,
	"Fr735xt":                    GarminProduct_Fr735xt,
	"VivoActiveApac":             GarminProduct_VivoActiveApac,
	"Vector2":                    GarminProduct_Vector2,
	"Vector2s":                   GarminProduct_Vector2s,
	"Virbxe":                     GarminProduct_Virbxe,
	"Fr620Taiwan":                GarminProduct_Fr620Taiwan,
	"Fr220Taiwan":                GarminProduct_Fr220Taiwan,
	"Truswing":                   GarminProduct_Truswing,
	"D2airvenu":                  GarminProduct_D2airvenu,
	"Fenix3China":                GarminProduct_Fenix3China,
	"Fenix3Twn":                  GarminProduct_Fenix3Twn,
	"VariaHeadlight":             GarminProduct_VariaHeadlight,
	"VariaTaillightOld":          GarminProduct_VariaTaillightOld,
	"EdgeExplore1000":            GarminProduct_EdgeExplore1000,
	"Fr225Asia":                  GarminProduct_Fr225Asia,
	"VariaRadarTaillight":        GarminProduct_VariaRadarTaillight,
	"VariaRadarDisplay":          GarminProduct_VariaRadarDisplay,
	"Edge20":                     GarminProduct_Edge20,
	"Edge520Asia":                GarminProduct_Edge520Asia,
	"Edge520Japan":               GarminProduct_Edge520Japan,
	"D2Bravo":                    GarminProduct_D2Bravo,
	"ApproachS20":                GarminProduct_ApproachS20,
	"VivoSmart2":                 GarminProduct_VivoSmart2,
	"Edge1000Thai":               GarminProduct_Edge1000Thai,
	"VariaRemote":                GarminProduct_VariaRemote,
	"Edge25Asia":                 GarminProduct_Edge25Asia,
	"Edge25Jpn":                  GarminProduct_Edge25Jpn,
	"Edge20Asia":                 GarminProduct_Edge20Asia,
	"ApproachX40":                GarminProduct_ApproachX40,
	"Fenix3Japan":                GarminProduct_Fenix3Japan,
	"VivoSmartEmea":              GarminProduct_VivoSmartEmea,
	"Fr630Asia":                  GarminProduct_Fr630Asia,
	"Fr630Jpn":                   GarminProduct_Fr630Jpn,
	"Fr230Jpn":                   GarminProduct_Fr230Jpn,
	"Hrm4Run":                    GarminProduct_Hrm4Run,
	"EpixJapan":                  GarminProduct_EpixJapan,
	"VivoActiveHr":               GarminProduct_VivoActiveHr,
	"VivoSmartGpsHr":             GarminProduct_VivoSmartGpsHr,
	"VivoSmartHr":                GarminProduct_VivoSmartHr,
	"VivoSmartHrAsia":            GarminProduct_VivoSmartHrAsia,
	"VivoSmartGpsHrAsia":         GarminProduct_VivoSmartGpsHrAsia,
	"VivoMove":                   GarminProduct_VivoMove,
	"VariaTaillight":             GarminProduct_VariaTaillight,
	"Fr235Asia":                  GarminProduct_Fr235Asia,
	"Fr235Japan":                 GarminProduct_Fr235Japan,
	"VariaVision":                GarminProduct_VariaVision,
	"VivoFit3":                   GarminProduct_VivoFit3,
	"Fenix3Korea":                GarminProduct_Fenix3Korea,
	"Fenix3Sea":                  GarminProduct_Fenix3Sea,
	"Fenix3Hr":                   GarminProduct_Fenix3Hr,
	"VirbUltra30":                GarminProduct_VirbUltra30,
	"IndexSmartScale":            GarminProduct_IndexSmartScale,
	"Fr235":                      GarminProduct_Fr235,
	"Fenix3Chronos":              GarminProduct_Fenix3Chronos,
	"Oregon7xx":                  GarminProduct_Oregon7xx,
	"Rino7xx":                    GarminProduct_Rino7xx,
	"EpixKorea":                  GarminProduct_EpixKorea,
	"Fenix3HrChn":                GarminProduct_Fenix3HrChn,
	"Fenix3HrTwn":                GarminProduct_Fenix3HrTwn,
	"Fenix3HrJpn":                GarminProduct_Fenix3HrJpn,
	"Fenix3HrSea":                GarminProduct_Fenix3HrSea,
	"Fenix3HrKor":                GarminProduct_Fenix3HrKor,
	"Nautix":                     GarminProduct_Nautix,
	"VivoActiveHrApac":           GarminProduct_VivoActiveHrApac,
	"Fr35":                       GarminProduct_Fr35,
	"Oregon7xxWw":                GarminProduct_Oregon7xxWw,
	"Edge820":                    GarminProduct_Edge820,
	"EdgeExplore820":             GarminProduct_EdgeExplore820,
	"Fr735xtApac":                GarminProduct_Fr735xtApac,
	"Fr735xtJapan":               GarminProduct_Fr735xtJapan,
	"Fenix5s":                    GarminProduct_Fenix5s,
	"D2BravoTitanium":            GarminProduct_D2BravoTitanium,
	"VariaUt800":                 GarminProduct_VariaUt800,
	"RunningDynamicsPod":         GarminProduct_RunningDynamicsPod,
	"Edge820China":               GarminProduct_Edge820China,
	"Edge820Japan":               GarminProduct_Edge820Japan,
	"Fenix5x":                    GarminProduct_Fenix5x,
	"VivoFitJr":                  GarminProduct_VivoFitJr,
	"VivoSmart3":                 GarminProduct_VivoSmart3,
	"VivoSport":                  GarminProduct_VivoSport,
	"Edge820Taiwan":              GarminProduct_Edge820Taiwan,
	"Edge820Korea":               GarminProduct_Edge820Korea,
	"Edge820Sea":                 GarminProduct_Edge820Sea,
	"Fr35Hebrew":                 GarminProduct_Fr35Hebrew,
	"ApproachS60":                GarminProduct_ApproachS60,
	"Fr35Apac":                   GarminProduct_Fr35Apac,
	"Fr35Japan":                  GarminProduct_Fr35Japan,
	"Fenix3ChronosAsia":          GarminProduct_Fenix3ChronosAsia,
	"Virb360":                    GarminProduct_Virb360,
	"Fr935":                      GarminProduct_Fr935,
	"Fenix5":                     GarminProduct_Fenix5,
	"Vivoactive3":                GarminProduct_Vivoactive3,
	"Fr235ChinaNfc":              GarminProduct_Fr235ChinaNfc,
	"Foretrex601701":             GarminProduct_Foretrex601701,
	"VivoMoveHr":                 GarminProduct_VivoMoveHr,
	"Edge1030":                   GarminProduct_Edge1030,
	"Fr35Sea":                    GarminProduct_Fr35Sea,
	"Vector3":                    GarminProduct_Vector3,
	"Fenix5Asia":                 GarminProduct_Fenix5Asia,
	"Fenix5sAsia":                GarminProduct_Fenix5sAsia,
	"Fenix5xAsia":                GarminProduct_Fenix5xAsia,
	"ApproachZ80":                GarminProduct_ApproachZ80,
	"Fr35Korea":                  GarminProduct_Fr35Korea,
	"D2charlie":                  GarminProduct_D2charlie,
	"VivoSmart3Apac":             GarminProduct_VivoSmart3Apac,
	"VivoSportApac":              GarminProduct_VivoSportApac,
	"Fr935Asia":                  GarminProduct_Fr935Asia,
	"Descent":                    GarminProduct_Descent,
	"VivoFit4":                   GarminProduct_VivoFit4,
	"Fr645":                      GarminProduct_Fr645,
	"Fr645m":                     GarminProduct_Fr645m,
	"Fr30":                       GarminProduct_Fr30,
	"Fenix5sPlus":                GarminProduct_Fenix5sPlus,
	"Edge130":                    GarminProduct_Edge130,
	"Edge1030Asia":               GarminProduct_Edge1030Asia,
	"Vivosmart4":                 GarminProduct_Vivosmart4,
	"VivoMoveHrAsia":             GarminProduct_VivoMoveHrAsia,
	"ApproachX10":                GarminProduct_ApproachX10,
	"Fr30Asia":                   GarminProduct_Fr30Asia,
	"Vivoactive3mW":              GarminProduct_Vivoactive3mW,
	"Fr645Asia":                  GarminProduct_Fr645Asia,
	"Fr645mAsia":                 GarminProduct_Fr645mAsia,
	"EdgeExplore":                GarminProduct_EdgeExplore,
	"Gpsmap66":                   GarminProduct_Gpsmap66,
	"ApproachS10":                GarminProduct_ApproachS10,
	"Vivoactive3mL":              GarminProduct_Vivoactive3mL,
	"ApproachG80":                GarminProduct_ApproachG80,
	"Edge130Asia":                GarminProduct_Edge130Asia,
	"Edge1030Bontrager":          GarminProduct_Edge1030Bontrager,
	"Fenix5Plus":                 GarminProduct_Fenix5Plus,
	"Fenix5xPlus":                GarminProduct_Fenix5xPlus,
	"Edge520Plus":                GarminProduct_Edge520Plus,
	"Fr945":                      GarminProduct_Fr945,
	"Edge530":                    GarminProduct_Edge530,
	"Edge830":                    GarminProduct_Edge830,
	"InstinctEsports":            GarminProduct_InstinctEsports,
	"Fenix5sPlusApac":            GarminProduct_Fenix5sPlusApac,
	"Fenix5xPlusApac":            GarminProduct_Fenix5xPlusApac,
	"Edge520PlusApac":            GarminProduct_Edge520PlusApac,
	"DescentT1":                  GarminProduct_DescentT1,
	"Fr235lAsia":                 GarminProduct_Fr235lAsia,
	"Fr245Asia":                  GarminProduct_Fr245Asia,
	"VivoActive3mApac":           GarminProduct_VivoActive3mApac,
	"Gen3Bsm":                    GarminProduct_Gen3Bsm,
	"Gen3Bcm":                    GarminProduct_Gen3Bcm,
	"VivoSmart4Asia":             GarminProduct_VivoSmart4Asia,
	"Vivoactive4Small":           GarminProduct_Vivoactive4Small,
	"Vivoactive4Large":           GarminProduct_Vivoactive4Large,
	"Venu":                       GarminProduct_Venu,
	"MarqDriver":                 GarminProduct_MarqDriver,
	"MarqAviator":                GarminProduct_MarqAviator,
	"MarqCaptain":                GarminProduct_MarqCaptain,
	"MarqCommander":              GarminProduct_MarqCommander,
	"MarqExpedition":             GarminProduct_MarqExpedition,
	"MarqAthlete":                GarminProduct_MarqAthlete,
	"DescentMk2":                 GarminProduct_DescentMk2,
	"Gpsmap66i":                  GarminProduct_Gpsmap66i,
	"Fenix6SSport":               GarminProduct_Fenix6SSport,
	"Fenix6S":                    GarminProduct_Fenix6S,
	"Fenix6Sport":                GarminProduct_Fenix6Sport,
	"Fenix6":                     GarminProduct_Fenix6,
	"Fenix6x":                    GarminProduct_Fenix6x,
	"HrmDual":                    GarminProduct_HrmDual,
	"HrmPro":                     GarminProduct_HrmPro,
	"VivoMove3Premium":           GarminProduct_VivoMove3Premium,
	"ApproachS40":                GarminProduct_ApproachS40,
	"Fr245mAsia":                 GarminProduct_Fr245mAsia,
	"Edge530Apac":                GarminProduct_Edge530Apac,
	"Edge830Apac":                GarminProduct_Edge830Apac,
	"VivoMove3":                  GarminProduct_VivoMove3,
	"VivoActive4SmallAsia":       GarminProduct_VivoActive4SmallAsia,
	"VivoActive4LargeAsia":       GarminProduct_VivoActive4LargeAsia,
	"VivoActive4OledAsia":        GarminProduct_VivoActive4OledAsia,
	"Swim2":                      GarminProduct_Swim2,
	"MarqDriverAsia":             GarminProduct_MarqDriverAsia,
	"MarqAviatorAsia":            GarminProduct_MarqAviatorAsia,
	"VivoMove3Asia":              GarminProduct_VivoMove3Asia,
	"Fr945Asia":                  GarminProduct_Fr945Asia,
	"VivoActive3tChn":            GarminProduct_VivoActive3tChn,
	"MarqCaptainAsia":            GarminProduct_MarqCaptainAsia,
	"MarqCommanderAsia":          GarminProduct_MarqCommanderAsia,
	"MarqExpeditionAsia":         GarminProduct_MarqExpeditionAsia,
	"MarqAthleteAsia":            GarminProduct_MarqAthleteAsia,
	"IndexSmartScale2":           GarminProduct_IndexSmartScale2,
	"InstinctSolar":              GarminProduct_InstinctSolar,
	"Fr45Asia":                   GarminProduct_Fr45Asia,
	"Vivoactive3Daimler":         GarminProduct_Vivoactive3Daimler,
	"LegacyRey":                  GarminProduct_LegacyRey,
	"LegacyDarthVader":           GarminProduct_LegacyDarthVader,
	"LegacyCaptainMarvel":        GarminProduct_LegacyCaptainMarvel,
	"LegacyFirstAvenger":         GarminProduct_LegacyFirstAvenger,
	"Fenix6sSportAsia":           GarminProduct_Fenix6sSportAsia,
	"Fenix6sAsia":                GarminProduct_Fenix6sAsia,
	"Fenix6SportAsia":            GarminProduct_Fenix6SportAsia,
	"Fenix6Asia":                 GarminProduct_Fenix6Asia,
	"Fenix6xAsia":                GarminProduct_Fenix6xAsia,
	"LegacyCaptainMarvelAsia":    GarminProduct_LegacyCaptainMarvelAsia,
	"LegacyFirstAvengerAsia":     GarminProduct_LegacyFirstAvengerAsia,
	"LegacyReyAsia":              GarminProduct_LegacyReyAsia,
	"LegacyDarthVaderAsia":       GarminProduct_LegacyDarthVaderAsia,
	"DescentMk2s":                GarminProduct_DescentMk2s,
	"Edge130Plus":                GarminProduct_Edge130Plus,
	"Edge1030Plus":               GarminProduct_Edge1030Plus,
	"Rally200":                   GarminProduct_Rally200,
	"Fr745":                      GarminProduct_Fr745,
	"Venusq":                     GarminProduct_Venusq,
	"Lily":                       GarminProduct_Lily,
	"MarqAdventurer":             GarminProduct_MarqAdventurer,
	"Enduro":                     GarminProduct_Enduro,
	"Swim2Apac":                  GarminProduct_Swim2Apac,
	"MarqAdventurerAsia":         GarminProduct_MarqAdventurerAsia,
	"Fr945Lte":                   GarminProduct_Fr945Lte,
	"DescentMk2Asia":             GarminProduct_DescentMk2Asia,
	"Venu2":                      GarminProduct_Venu2,
	"Venu2s":                     GarminProduct_Venu2s,
	"VenuDaimlerAsia":            GarminProduct_VenuDaimlerAsia,
	"MarqGolfer":                 GarminProduct_MarqGolfer,
	"VenuDaimler":                GarminProduct_VenuDaimler,
	"Fr745Asia":                  GarminProduct_Fr745Asia,
	"VariaRct715":                GarminProduct_VariaRct715,
	"LilyAsia":                   GarminProduct_LilyAsia,
	"Edge1030PlusAsia":           GarminProduct_Edge1030PlusAsia,
	"Edge130PlusAsia":            GarminProduct_Edge130PlusAsia,
	"ApproachS12":                GarminProduct_ApproachS12,
	"EnduroAsia":                 GarminProduct_EnduroAsia,
	"VenusqAsia":                 GarminProduct_VenusqAsia,
	"Edge1040":                   GarminProduct_Edge1040,
	"MarqGolferAsia":             GarminProduct_MarqGolferAsia,
	"Venu2Plus":                  GarminProduct_Venu2Plus,
	"Gnss":                       GarminProduct_Gnss,
	"Fr55":                       GarminProduct_Fr55,
	"Instinct2":                  GarminProduct_Instinct2,
	"Fenix7s":                    GarminProduct_Fenix7s,
	"Fenix7":                     GarminProduct_Fenix7,
	"Fenix7x":                    GarminProduct_Fenix7x,
	"Fenix7sApac":                GarminProduct_Fenix7sApac,
	"Fenix7Apac":                 GarminProduct_Fenix7Apac,
	"Fenix7xApac":                GarminProduct_Fenix7xApac,
	"ApproachG12":                GarminProduct_ApproachG12,
	"DescentMk2sAsia":            GarminProduct_DescentMk2sAsia,
	"ApproachS42":                GarminProduct_ApproachS42,
	"EpixGen2":                   GarminProduct_EpixGen2,
	"EpixGen2Apac":               GarminProduct_EpixGen2Apac,
	"Venu2sAsia":                 GarminProduct_Venu2sAsia,
	"Venu2Asia":                  GarminProduct_Venu2Asia,
	"Fr945LteAsia":               GarminProduct_Fr945LteAsia,
	"VivoMoveSport":              GarminProduct_VivoMoveSport,
	"VivomoveTrend":              GarminProduct_VivomoveTrend,
	"ApproachS12Asia":            GarminProduct_ApproachS12Asia,
	"Fr255Music":                 GarminProduct_Fr255Music,
	"Fr255SmallMusic":            GarminProduct_Fr255SmallMusic,
	"Fr255":                      GarminProduct_Fr255,
	"Fr255Small":                 GarminProduct_Fr255Small,
	"ApproachG12Asia":            GarminProduct_ApproachG12Asia,
	"ApproachS42Asia":            GarminProduct_ApproachS42Asia,
	"DescentG1":                  GarminProduct_DescentG1,
	"Venu2PlusAsia":              GarminProduct_Venu2PlusAsia,
	"Fr955":                      GarminProduct_Fr955,
	"Fr55Asia":                   GarminProduct_Fr55Asia,
	"Edge540":                    GarminProduct_Edge540,
	"Edge840":                    GarminProduct_Edge840,
	"Vivosmart5":                 GarminProduct_Vivosmart5,
	"Instinct2Asia":              GarminProduct_Instinct2Asia,
	"MarqGen2":                   GarminProduct_MarqGen2,
	"Venusq2":                    GarminProduct_Venusq2,
	"Venusq2music":               GarminProduct_Venusq2music,
	"MarqGen2Aviator":            GarminProduct_MarqGen2Aviator,
	"D2AirX10":                   GarminProduct_D2AirX10,
	"HrmProPlus":                 GarminProduct_HrmProPlus,
	"DescentG1Asia":              GarminProduct_DescentG1Asia,
	"Tactix7":                    GarminProduct_Tactix7,
	"InstinctCrossover":          GarminProduct_InstinctCrossover,
	"EdgeExplore2":               GarminProduct_EdgeExplore2,
	"DescentMk3":                 GarminProduct_DescentMk3,
	"DescentMk3i":                GarminProduct_DescentMk3i,
	"ApproachS70":                GarminProduct_ApproachS70,
	"Fr265Large":                 GarminProduct_Fr265Large,
	"Fr265Small":                 GarminProduct_Fr265Small,
	"Venu3":                      GarminProduct_Venu3,
	"Venu3s":                     GarminProduct_Venu3s,
	"TacxNeoSmart":               GarminProduct_TacxNeoSmart,
	"TacxNeo2Smart":              GarminProduct_TacxNeo2Smart,
	"TacxNeo2TSmart":             GarminProduct_TacxNeo2TSmart,
	"TacxNeoSmartBike":           GarminProduct_TacxNeoSmartBike,
	"TacxSatoriSmart":            GarminProduct_TacxSatoriSmart,
	"TacxFlowSmart":              GarminProduct_TacxFlowSmart,
	"TacxVortexSmart":            GarminProduct_TacxVortexSmart,
	"TacxBushidoSmart":           GarminProduct_TacxBushidoSmart,
	"TacxGeniusSmart":            GarminProduct_TacxGeniusSmart,
	"TacxFluxFluxSSmart":         GarminProduct_TacxFluxFluxSSmart,
	"TacxFlux2Smart":             GarminProduct_TacxFlux2Smart,
	"TacxMagnum":                 GarminProduct_TacxMagnum,
	"Edge1040Asia":               GarminProduct_Edge1040Asia,
	"EpixGen2Pro42":              GarminProduct_EpixGen2Pro42,
	"EpixGen2Pro47":              GarminProduct_EpixGen2Pro47,
	"EpixGen2Pro51":              GarminProduct_EpixGen2Pro51,
	"Fr965":                      GarminProduct_Fr965,
	"Enduro2":                    GarminProduct_Enduro2,
	"Fenix7sProSolar":            GarminProduct_Fenix7sProSolar,
	"Fenix7ProSolar":             GarminProduct_Fenix7ProSolar,
	"Fenix7xProSolar":            GarminProduct_Fenix7xProSolar,
	"Lily2":                      GarminProduct_Lily2,
	"Instinct2x":                 GarminProduct_Instinct2x,
	"Vivoactive5":                GarminProduct_Vivoactive5,
	"Fr165":                      GarminProduct_Fr165,
	"Fr165Music":                 GarminProduct_Fr165Music,
	"Edge1050":                   GarminProduct_Edge1050,
	"DescentT2":                  GarminProduct_DescentT2,
	"HrmFit":                     GarminProduct_HrmFit,
	"MarqGen2Commander":          GarminProduct_MarqGen2Commander,
	"LilyAthlete":                GarminProduct_LilyAthlete,
	"Fenix8Solar":                GarminProduct_Fenix8Solar,
	"Fenix8SolarLarge":           GarminProduct_Fenix8SolarLarge,
	"Fenix8Small":                GarminProduct_Fenix8Small,
	"Fenix8":                     GarminProduct_Fenix8,
	"D2Mach1Pro":                 GarminProduct_D2Mach1Pro,
	"Enduro3":                    GarminProduct_Enduro3,
	"FenixE":                     GarminProduct_FenixE,
	"Sdm4":                       GarminProduct_Sdm4,
	"EdgeRemote":                 GarminProduct_EdgeRemote,
	"TacxTrainingAppWin":         GarminProduct_TacxTrainingAppWin,
	"TacxTrainingAppMac":         GarminProduct_TacxTrainingAppMac,
	"TacxTrainingAppMacCatalyst": GarminProduct_TacxTrainingAppMacCatalyst,
	"TrainingCenter":             GarminProduct_TrainingCenter,
	"TacxTrainingAppAndroid":     GarminProduct_TacxTrainingAppAndroid,
	"TacxTrainingAppIos":         GarminProduct_TacxTrainingAppIos,
	"TacxTrainingAppLegacy":      GarminProduct_TacxTrainingAppLegacy,
	"ConnectiqSimulator":         GarminProduct_ConnectiqSimulator,
	"AndroidAntplusPlugin":       GarminProduct_AndroidAntplusPlugin,
	"Connect":                    GarminProduct_Connect,
}

func GarminProductFromString(s string) GarminProduct {
	if v, ok := GarminProductValues[s]; ok {
		return v
	}
	return GarminProduct(GarminProduct_Invalid)
}
