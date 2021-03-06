package vkgtaro

var (
	photo = []string{
		"                                                                                                                                                ",
		"                                                      ..gMMNNN#NNNgg-..                                                                         ",
		"                                                    .dMMNNNNNNMMNNNNMN#MN&,.                                                                    ",
		"                                                  .dMNNMMMMMNNNNNNNNNMMNNNkqm-...                                                               ",
		"                                               .+MNMNNN##NNNNMNMMNNNMNMMMMNNNMMHHHm&,                                                           ",
		"                                            .(MNMNNNMNNNNNMMNMMMMNNMNMNMNMNNNNNNMMMNNm&.                                                        ",
		"                                          .dMNMMNMNNNMMMMMNNMMNMNNMNM#MMMMMMNNNNMMMMMMMN,                                                       ",
		"                                        .dNMNMMNMMNNNNNN##NMNNMMMNMMMHHHHMMNMNNM#NNNMNNMN;                                                      ",
		"                                       .MMMMMNNMNNNNNNN#NM#MMMMMMMHHWWkwXHHMMMMNNMNNMNMNNN-                                                     ",
		"                                      (MNNNNMNMNNMNNNMN#MMMMgHHHWWWXXwZzXXWHMMMMNNNNMNNMNNb                                                     ",
		"                                     .MMNMNMNNMMNNNNNMMHHmqHHHWWX00tOrZOttXWHMM#NNNMMMNNNMM[                                                    ",
		"                                    .NMNNNMMMNNNNNNNMMMMHHHW0UXzOOI?+11zlzrXWH@M#NNNMNMNNNNN.                                                   ",
		"                                    JMNMMNMMMMNMM##MMMMHW0XZOOI1====<><+?1zwXWHMMNNNMNMNNNMM]                                                   ",
		"                                   .MNNMMMMNNMNMHMMMNMHHW0wlzz?=><>?><>+??1zwXW@MNNNNNNMNNNNN                                                   ",
		"                                   ,MNMNNNMNNNMMHMHHHHWWW0tll1???????>???>?1=OXWMNNNNNNNNNNN#.                                                  ",
		"                                   ,MMMMMNMNN#MHWWHHfWUXSZv=1??>11===?==zlzzllOXWMMNNNNNNMNNM{                                                  ",
		"                                   ,NNNMMNMMMMHfXWWXZwwOO1??>?+??=???=zOwwwOOttwXWMNNNMNNNNNM}                                                  ",
		"                                   ,MMNMMNNMMHffWkkkAAwOzz1??=??=zzuQQQkkWyfXXrrrwWMMNNMNNNNM:                                                  ",
		"                                   ,NMMMMMMMHpWWgg@@@M@NmAwOOz1zwXWMH9UZOwUXXVXOwwwWMMNMNNNNM                                                   ",
		"                                   .MNNMMMMHppkHHWXZOvzwUWWXwv=zZUUwOOO=?1twwrttOrtwWMNNNNNNN                                                   ",
		"                                    MMMMNMMpppppWXXOzzzzzwwWkv1zlrXZXWNHHmQykrOlzOOwXNMMNMMNF                                                   ",
		"                                    ,MMNMMNkWKWWWWHMMMMWHXkXHZvvSX0UWWUUUOOUX0lVOTwZVMNMMMNMt                                                   ",
		"                                     4NMMMMpWXZWWHWUUzuZXuXySO=1OrvvzOOOVOvCI=z1?=zrwMN#@MN#                                                    ",
		"                                      MMNMHfyuuXu0XU0VOtwXyyuOz1lOOOO====?1z?????=zzXMMM@M#F                                                    ",
		"                                      .MNNNpyuuvrtOlzOtrvXyXVz=11zttlllll=?<><?===twXMM0H#M`                                                    ",
		"                                       ,MNMHyyZXtOl=OOOwXXXXO11?>1z=lOwOlz?<;(<?==trwHMSd#D                                                     ",
		"                                        ,MHHWyXXrl==1=OwXkXXkwwwOAAszI1OXwO=+<<?=lltdHH0d@`                                                     ",
		"                                         (HNVVWXwOlzzOwuZWqHHWUUwrwVI=+?1VuwO==?=lltwHRQM!                                                      ",
		"                                          ,MWVVWXrOltwXuXXXXZOOwOv1zlz1z+?1OZ<?1?=zOwWHM|                                                       ",
		"                                           WHyyZuvtlrrvuZuzZlz1???????ztOwO+O<;>>?=ldM@$                                                        ",
		"                                           .HWZXuwIzrOXWyXXXQkkwZTXVXwAQmXkzz<<<;1=ldMF                                                         ",
		"                                            ,XyuuvOzZOWWQNHHX$~_~-((zsX9111?<<>;>+zwM#                                                          ",
		"                                             .WXwuzO=<zOXWHHWWU0UV0OwOC?;;<>;?>??zOdM^                                                          ",
		"                                               TXwXXOzzzwuXXUXXrrOOz>+><;;>1<+?+zOdMt                                                           ",
		"                                                ?XXXywtllOXuXzzzvrOz<<;;:;;+1?1ztwMD                                                            ",
		"                                                  7yyykOlz1ztOlzz1?>;::;;:+1=zzrOt=                                                             ",
		"                                                   .4XyWXOz>???<>>;;;::;;+=zzrOO==_                                                             ",
		"                                                     WyyyXXwz+1z>+++?+?+zltttlz???u,                                                            ",
		"                                                     .yyyyyyyXwzzzlzzltttllz1z<>>?zMN,                                                          ",
		"                                                      XyyyyyyZyXuuwOOtOl=?=?<;;;>>1dMMm.                                                        ",
		"                                                     .dZZZyZyyZZXZOOOz1?>><;;;;;;>>zdNNHe.                                                      ",
		"                                                     JNyZZZyZZZuwlz?<;;;;;:::::::;;1ZM#NMNm,                                                    ",
		"                                                    .MMkWyyyZZZZXOz?<:;;::::::::::;+zHH#HMNNMHk>_                                               ",
		"                                                   .MMMHfyZZZZZuuuOz>;::::::~:::::;;dMNMMH#N###Hx_..  ..                                        ",
		"                                                  .HMMMHWyyuuuuuzwOz<;:::~~~~~:~~::+dMMHHHM#####Nn_..........~..-                               ",
		"                   ++>-.                         .dHMMMqpyyykrOwwOl?>;::~~~~~:~~~:(d#N@CwWH###H##Ns_...................._                       ",
		"                   .OOz><.                      .XWMMMKUWyyyXwOlz=?>;;::::~~:~~~_~<WNM><+dWH###H#HHc_.......................                    ",
		"                     (1z<<<:        `  ..-___-..WWMMMMRI=z7TU0O=???><<<<~_````.`` (M#D~_<zWMH#####Mk--.................````...                  ",
		"                       .z<<_~(<~~_~~~~~~~~:~_:+WNHWMHM$<<<<<;::~__~_..``````````.(WH#_. _+WH#HHH###N$. ......~.......~.````````.                ",
		"                       ._1<(_~~~_......~~~~~~(jMVSTMMH$_~_~__~~....`.``.`.`````. (MM%`.`_(dWHH#H#HHNW-.......~_.``.`-_`.``.```..`               ",
		"                `    .....1+:__~~__.....~~~~_(dMH8VlvT6:~___-(J+mgJ--&; `..`..``(HH@_``.._zWMH####HHH{_.....~~.```..._``.````...`.              ",
		"            ((((<-..__.....(1__~~~_-.`...((J__dMwO++zz+<_.._JMM@HgHHHHHH+```.`.-dMM>``.`.-jWM#######NR_..`._~.````....``.````...``              ",
		"           ?1+1+???<:_.`.~.. <_~~_~<-.`(dMHMm-(HHkAzz+j2_.-dM#MM@@@@HHWMN/` `` (2(3.```.` (XH##MMMM#NW-.`..~..````...``.``.`. ..```             ",
		"             _7Ozz<;<:(<___  .~~~~~~(z&(MHMMmmmQMMMHkQkHHMMMMMHWqHqkkkH#MHJ&J-(vV9! .```.._dWMMi+gJW#Hr.``.~..```....` .```` ..``..             ",
		"                _?1z++<<::<_..~~~~~~(OXX@HN#MHMNMMHMgHMMMM###M+X+__<(?zM#MMM#NNHHC` ``.```-(WMMmMN0d#Nk_`....`````..``..```...``````            ",
		"                ~~~~?C1+<:::<_~~:((+zdWHM#HHMMMNHMM@H@MMMNNMMMMMMMMMNMMM#NNMHMMMMo_.`.`.`. (XWHbjJvHHMW;`..._````...``_```. ...``.`..           ",
		"               _.~_____ ?Cy+__~:++zwXpW@M###MHNMMMMHHHMNMNNNNMMMMMNMMMNMMHNMHHbpHM{ .```.`._zW#N&zVTM#HI.....```.``.``````... ````..`           ",
		"               ..___.~_. __7W<~(jwwXWHHH##MMMMHHM#MMMMMMNNMMNNMMMMMMN#####N#HPUVXMQ#4e..``.-(WM#KYT(##NR_`...```..```````..._``.`..``.          ",
		"               ..._~~__`._(JUI<;+XVpWkMM###NNNNNMNHN#?<jMM@MM@HgmHqHHMMHM9HMMHHHWMYHAJQx..`.(yM#M6ucHMHH{...```..`.````...._````....`.          ",
		"              _.`.._~.~_._?ykOz::<?TWWH#MMMMMM#NNNMUC+>dHMHMMHgHqqHkWHHMk<_?BMRSH@} . _WN+.`_dW#Nx+JHM#HI.`.```.`.````.`.._````....``.-         ",
		"             ._.`..~~~~~.__ZWkwz+<<(+<+OTW@ggHMNMMHZ11dHHMMMMMMMMMMMHHWHHs<_<<7WHM>..`.(MHN, jVMHk+J(MMNR_.````...```.`...````.....``..         ",
		"            .:_..`..~~~_.~(wzXHHkkwyOOz+<<dHHMNMMM0zzdHHM#MNM#N#MNNNNMMMHHk<___~(H~.` .`?MMMe(VM#NJgg#MMHl.````..````....````......``.          ",
		"            ~:_.....~~~~__~CwXUwZUVWHkkXOzdHHNNMMNZ=OWMH##MMMNMMMMMM#HMMMHHk+++(+S_..`.`.(MM#NdWM#M#####NR_```....``.`...``.``.....``.          ",
		"           _:~:~~~.~~~~~~~~<?WWWkOzzZOOUUTMH##MMN#ZzzdH##HH@MNNMMM##MMNNNMHMS=??jC. .``..` ?MMMHM##NNMMMMMp.`` ..```....``.`....~.``..`         ",
		"          .~~::~~~~.~~~~~~~~(:?OVVkXkwwwzuWHHNNMNHZlOWHHMMMMNNMHXHMMMMMNNNHHI??>1!.....``..`.MHmHMHHH#H#HM#_`` ...``...`````..~...`` ```        ",
		"          ~~:~:~.~~.~..~~~~~<(:_?WXUUHHmmMMHHNMMN#ZlwMHHMMH@M##N##HMMH#NNMHRI?;>;_.......``.-(WM#MNM#HMMMMN~`` ...` ..``.`...~.~.```.`.`        ",
		"         .~~~~~~.~~~~.~~~~~~::::~zXXXWMN##MH#MNNMNkOzWHMHMMH@MNNNNMMMM#MNNMI?;;>?_.....`....~_<dMMMMHMHHHMD``` ...`..````.`.~~.~.`` ...`        ",
		"         ~~~:~~~~~.~~~.~~~~~~:::(wuyyZXWHMMMMMHyyWWXOXMHHHMHM@@@@HMHHMMNNMSzz>><><__.....`.___<+WM#####HMD_```.... _..``....~~_.```..``.        ",
		"       ` ~~~~~:~~~.~~~~~~~~~~~:<duXZuzrrzzuXXXMMMMNWkwXHM@MMHMMHHHM###M#M6Otv?>;;;:(dAwuu&&&&+&+++TM#M#N>..`.`-...__..``...~~~_````..``.        ",
		"        .~~~~~~:~~~~~~~~~~~~::<duuuuwwOOrzvwXWWpWMHyXOwzWH@@@MMMHHHHHMMY<<>>???>;;<<wh(_:~~<>1z17774kTMNNe.``._~.~_`.`...~~~~_`````.``.`        ",
		"        ~.~~~~~~_.~~~~~~~~~_~(wwzuuzvtttvzvXWWXWM0XXXXUyZuXWHHMMMMMBY1;<(<+???>;;:;:<Xk<::~:<<<~~:_(V -7MNNx.`_~~~.`....~~~~~``````.```.        ",
		"        .~~~~~~~:_~~~~~~~~::(wzwwZXOttltrzuXWuvwWHHkXOOwWyyXXwOz??++>1>???????<<<;::?U9?<<:::~~~~_JV_. `.TMMN,_~~_....~~~~~__`````..````.       ",
		"        ~~.~.~~~~::_~~:~~~~jzzrwZuvttlltrwXy0vtOOwXHMmOz=OXuZUXzzOtOzzzz==1z???>;::~:(?OQ&-_~~~~(JC..````.(WNNN-_....~~~__`````````.``.`.       ",
		"        ~~~~~.~~~~~~::~:~~+zzrrwZurtllltvzXSwZzlOvuXXXWWs<UKIXXVTTXXXuXwwwOOtz1+><::~~_<(UNn-~~(z:...```.`..?MM$~~~~~_..``````````.````..       ",
		"        _~~~~~~.~~~~~~::<jXrrrzuuwtlllltrwWW9WWAOwuurttzZWWNgdI>>zXWWBWkWXXXXuwOzz+<_:~~:(Wkko(Z!  .```..`...(><~~~~..````````````.`````.       ",
		"         :_~~~.~~~~.~~_(jXrtrvzzvrll=lttrXHD<;+ZWHkrOtlllzdWHgkiwXHHMHmSwXWVWZXwttlz+<::~(dWHmc~(--..`.``.-___~~~~_~_```.`````````..`.``.       ",
		"         ::~~~~~.~~_(jyuzrttrvzvrtlll=lrwWE<:<<+1zTHAz=lltwWNSWSQwwX#XWkrrwZUyyXXXwwOz<<(dyWWD_<>;;:---.._(<<~_~~~_~_````..```````.``````_      ",
		"         _:_~~~~~~(dUUuzOlltOvzrttl===ltw0<::;+&dkz<7Uk?=ldWMWWHkmwdWHHkmytrrQWNkWHHkmAdyyWWf~:<<::::<>;;<<<__.`` __~_````````````.``````       ",
		"          :~~~~~~(XXuzzrttlttrrrOll==tttZ>:(jdbpbHI<_Jf><1X0X#THrwwNmM@IvHytrXKMMMMNMNHkwXp6+<:::::::::;::;;<:?1+..``. ```````````.``.``..      ",
		"          :~::~(Jwwuuvrttttttrtttl=l==lwz(dyVVppbkI<<XI>;dR+OHgNQQd#N#C<;dRtwdKOVHMHH@H@MHWZO1<(::~~~~::::<:;::;<zXG+-.``````````.````.``.      ",
		"          ::~(XW0wzurttttrtttttt=====zwkWXZyyVpppHI<~dI;>zR++dNNRrdNMD<;;dSOOd0lltZHMMMMMHyXOzO><::::~::~~::<;:::::?zVXw&-.` ````.`.``````      ",
		"          ((dS0ruuuzrtttrrtttttl==?1<1ZTXuuZZyfpbHz<_(I(<<XWQXMNNwMMM&&+d9lzOWz??11zWH@HkyyZXuwOz+<_:~~:~~~~::~~~:~::(<?OwzX&-. ``````````      ",
		"      `   (0Xvvzuvzrttttttttlll=??>>jV~_JXuuuZyfbR<__.jc_:(+zdM#NW##@dRz1z11dD<<;;;<+WHH@HkWyZZuuwOz+<(:~~~~~:~~~~~~~::~~<1OZXfn.`````````      ",
		"      .<(JHzzzzuzrOllltttttlll==?>?j0! JzXzzzXWppC~_.` j+_:<1vMHMMHMmWCz><jd3<<~~<:::<WHHHMHkyXZXuuzwOzz+___~~::~~:~~~~~:~~;;1OXWo. ``````.     ",
		"      .dHyXZuwuuzOllllrtttll====?1uW>_Jvrvzwu0<?T<_.``.`?w+(<1OHNMH@v<<<jwY~.__...___~?WfWbHHHkZZZZuuwOrOzzz+<<<::_::~:~:::~;<<?zXXn-. ```.     ",
		"     .WfpyXZuuuzrO=l=ttrtll=l===zdW3_JwrvrvZf:~::<_...``` ?TXAydMHH$<jdV=~..._~..--_` _?WffWpkHkWXXXuuuzwrrrrOlzzz+<<<<;:;:;:;;>??+zOXWk&-.     ",
		"     .yVVVyyZuuvtll==ttrtllllllwwWD_JvvvzvzV_~:::~...`..`..-_dSdMH#<J>__..~.._(JwuX_.`..(XXZWVWHIOWXzuzXurrrrttttttOzzz1?++?>;;>;>;<>zyZuX}     ",
		"      ?UXWfWZuzOlll==ltltllltOwXWK<zzvvvvzuXAaaga&+J(-((.`..~d0dMHD_d!..~._(Juuuzzu-....-(vvwXXWWc?OZXXvrvzzwrttllllzOzzz===??>;>;;;>z0wC!      ",
		"         ?4Wyuvrtl==zlttlltwwXWpk3zzzzzzuzuZyWpbpWuzzzzu>....dIz<I~_w_`-(dZZXuuzzzvwOrwwrrvrrrrzuk-_<+zOrrtrtOOtttl===???><;<;;;<:;<+?!         ",
		"           Oyuzrll===?zltOwXWWbHD_zrvzzzzuuyfpbbpWuzzzzu:``._X=t_O_(X  XuZuuuzzzzzvrrrrrrrrrrvvrzuI._~<<+1OttOtOtlz==??>><;;;::;::<<:           ",
		"           .ZXrrtlll===zwXZyWH9j:`.7wzzuuuZXppbpfXuuuuzu!``..W=dMNc(0.` ?TZXuzzuzvrrtttrrtrrtrrvwX>`.._~:<<<1wwrtrttll==+>>;;;;:::+c            ",
		"            XuzrttllzltwwyVpk3(>_```......~~::<<~~__~~~~``. .R1vTH=($ ``.._?CvvvvzZ<<??77<7777CC7!.``..`__~~<?zOOtrrOOtz==1??+1+zzc             ",
		"            .XXuvrttOrwXyWWK>:+~.``````...~~:::~~...........($(+-<<({.``````` ?7vzr`````````````.`..```.``..__<<1Oz<?OwvrrOOllltO>              ",
		"              TWkXwrwzuZWp9><+<_.``.``....~:~~:~~......jI<?<?<<1w>~(7777Ol ` ``` ?~````````````````````.````.`._~<?<.._?wzzvOOrZ!               ",
		"               .4WXXXZZXX3:<<<_.```.``.`.~~~~:~~~.....`-j+~~~:;;+I~_~~_(C``.````````````````````````.```.`.`.``..-<<_..~_?OXXV^                 ",
		"                                                                                                                                                ",
	}

	characters = []string{
		"         ##                                                   ",
		"          #                  #                                ",
		"          #                  #                                ",
		"          #                  #                                ",
		"##   ##   #    #    ### ## ######     ####   ### ##     ###   ",
		"#     #   #   #    #   #     #       #    #    ##  #   #   #  ",
		"#     #   #  #     #   #     #            #    #   #  #     # ",
		" #   #    # ##     #   #     #        #####    #      #     # ",
		" #   #    ##  #     ###      #       #    #    #      #     # ",
		" #   #    #   #    #         #      #     #    #      #     # ",
		"  # #     #    #   ####      #   #  #     #    #      #     # ",
		"  # #     #    #  #    #     #   #  #    ##    #       #   #  ",
		"   #     ###   ## #     #     ###    #### ## ######     ###   ",
		"                  #     #                                     ",
		"                   #####                                      ",
	}
)
