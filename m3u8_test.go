package parsem3u8

import (
	"bytes"
	"testing"
)

// got from https://github.com/iptv-org/iptv/blob/master/channels/tr.m3u
var M3uData = bytes.NewBufferString(`#EXTM3U
#EXTINF:-1 tvg-id="AfyonTurkTV.tr" tvg-name="Afyon Turk TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/tXbD8Sz/8A1HOMZ.png" group-title="",Afyon Turk TV (720p)
https://5be5d840359c6.streamlock.net/afyonturktv/afyonturktv/playlist.m3u8
#EXTINF:-1 tvg-id="AhiTVKirsehir.tr" tvg-name="Ahi TV Kirsehir" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/yBGZy4n/obCp0aK.png" group-title="",Ahi TV Kirsehir
http://yayin3.canlitv.com:1935/canlitv/ahitv/chunklist.m3u8
#EXTINF:-1 tvg-id="AkitTV.tr" tvg-name="Akit TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/zG7NM55/Jagn7PG.png" group-title="",Akit TV (720p)
https://akittv-live.ercdn.net/akittv/akittv.m3u8
#EXTINF:-1 tvg-id="AlZahraTVTurkic.tr" tvg-name="Al-Zahra TV Turkic" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/1LGHMzQ/lwkowCS.png" group-title="",Al-Zahra TV Turkic (1080p)
https://live.al-zahratv.com/live/playlist.m3u8
#EXTINF:-1 tvg-id="AlZahraTVTurkic.tr" tvg-name="Al-Zahra TV Turkic" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/1LGHMzQ/lwkowCS.png" group-title="",Al-Zahra TV Turkic
https://live.al-zahratv.com/live/playlist2/index.m3u8
#EXTINF:-1 tvg-id="AltasTV.tr" tvg-name="Altaş TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/37VNK2B/Nnsnayw.png" group-title="",Altaş TV (720p)
http://stream.taksimbilisim.com:1935/altastv/bant1/playlist.m3u8
#EXTINF:-1 tvg-id="BesiktasWebTV.tr" tvg-name="Beşiktaş Web TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/bRhnM6W/qti5aea.png" group-title="",Beşiktaş Web TV
https://s01.vpis.io/besiktas/besiktas.m3u8
#EXTINF:-1 tvg-id="BloombergHT.tr" tvg-name="Bloomberg HT" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://www.bloomberght.com/assets/v2/images/desktop/logo/bloomberght_logo.svg" group-title="Business",Bloomberg HT (720p)
https://ciner.daioncdn.net/bloomberght/bloomberght.m3u8
#EXTINF:-1 tvg-id="BRT1Kibris.tr" tvg-name="BRT 1 (Kıbrıs)" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/BBx9nKK/yCer0tL.png" group-title="",BRT 1 (Kıbrıs)
http://wms.brtk.net:1935/live/brt1/playlist.m3u8
#EXTINF:-1 tvg-id="BRT2.tr" tvg-name="BRT 2" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/0Fz0rSB/yoNkkLc.png" group-title="",BRT 2
http://wms.brtk.net:1935/live/brt2/BratuMarian.m3u8
#EXTINF:-1 tvg-id="BRT2Turkish.tr" tvg-name="BRT 2 (Turkish)" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/0Fz0rSB/yoNkkLc.png" group-title="",BRT 2 (Turkish)
http://wms.brtk.net:1935/live/brt2/playlist.m3u8
#EXTINF:-1 tvg-id="BursaTV.tr" tvg-name="Bursa TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/N2Yc8Yf/k6DUajb.png" group-title="",Bursa TV
https://win1.yayin.com.tr/bursatv/bursatv/chunklist.m3u8
#EXTINF:-1 tvg-id="CanTV.tr" tvg-name="Can TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/qdBr2fH/kgwc6oT.jpg" group-title="",Can TV
http://canbroadcast.com:7000/canlican/tv.m3u8
#EXTINF:-1 tvg-id="CayTV.tr" tvg-name="Cay TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/80d40Qb/e2K5ixJ.png" group-title="",Cay TV (720p)
http://stream2.taksimbilisim.com:1935/caytv/bant1/CAYTV.m3u8
#EXTINF:-1 tvg-id="CRITurkBelgesel.tr" tvg-name="CRI Turk Belgesel" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/xXjjbq6/yhzxKnt.png" group-title="",CRI Turk Belgesel (480p)
http://cri.aa.net.tr:1935/belgesel/belgesel/playlist.m3u8
#EXTINF:-1 tvg-id="DogusTV.tr" tvg-name="Doğuş TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/Gdb8yX4/xyubJha.png" group-title="",Doğuş TV
https://s01.vpis.io/dogustv/dogustv.m3u8
#EXTINF:-1 tvg-id="ETVManisa.tr" tvg-name="ETV Manisa" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/QKXtdkz/sOd3INg.png" group-title="",ETV Manisa
https://broadcasttr.com:446/manisaetv/bant1/chunklist.m3u8
#EXTINF:-1 tvg-id="EuroStar.tr" tvg-name="EuroStar" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/N7gDm9P/7g6oP1l.jpg" group-title="",EuroStar
https://xrklj56s.rocketcdn.com/eurostar.stream_720p/chunklist.m3u8
#EXTINF:-1 tvg-id="FinansTurkTV.tr" tvg-name="Finans Turk TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/wBwmB1T/iY0osc7.png" group-title="",Finans Turk TV (360p)
http://live.arkumedia.com:1935/finansturktv/finansturktv/playlist.m3u8
#EXTINF:-1 tvg-id="HalkTV.tr" tvg-name="Halk TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.imgur.com/vYEnfbj.jpg" group-title="",Halk TV (720p)
https://mn-nl.mncdn.com/blutv_halktv/smil:halktv_sd.smil/playlist.m3u8
#EXTINF:-1 tvg-id="IBBTV.tr" tvg-name="IBB TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/d0T2jfg/ibb-tv-logo-white.png" group-title="",IBB TV
rtmp://wowza.istweb.tv:1935/webtv/webtv_wowza1
#EXTINF:-1 tvg-id="IstanbulMetropolitanMunicipality.tr" tvg-name="Istanbul Metropolitan Municipality" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/d0T2jfg/ibb-tv-logo-white.png" group-title="",Istanbul Metropolitan Municipality (720p)
http://wowza.istweb.tv:1935/dp/istanbul2/playlist.m3u8
#EXTINF:-1 tvg-id="Kanal3.tr" tvg-name="Kanal 3" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/Jq49Y1g/uHAGsRk.png" group-title="",Kanal 3 (720p)
http://stream.taksimbilisim.com:1935/kanal3/bant1/Kanal3.m3u8
#EXTINF:-1 tvg-id="Kanal7.tr" tvg-name="Kanal 7" tvg-country="TR" tvg-language="Turkish" tvg-logo="" group-title="",Kanal 7
https://live.kanal7.com/live/kanal7LiveDesktop/index.m3u8
#EXTINF:-1 tvg-id="Kanal7.tr" tvg-name="Kanal 7" tvg-country="TR" tvg-language="Turkish" tvg-logo="" group-title="",Kanal 7
https://live.kanal7.com/live/kanal7LiveMobile/index.m3u8
#EXTINF:-1 tvg-id="KanalD.tr" tvg-name="Kanal D" tvg-country="TR" tvg-language="Turkish" tvg-logo="" group-title="",Kanal D
https://stream1.kanald.ro/iphone/live.m3u8
#EXTINF:-1 tvg-id="KonTV.tr" tvg-name="Kon TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/KjVF97d/D8uqiZD.jpg" group-title="",Kon TV (1080p)
https://59cba4d34b678.streamlock.net/live/kontv/playlist.m3u8
#EXTINF:-1 tvg-id="Kral Pop TR" tvg-name="Kral Pop TR" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/7v8wmsb/7H8WKei.png" group-title="Music",Kral Pop
https://0ajuu84p.rocketcdn.com/kralpop_720/chunklist.m3u8
#EXTINF:-1 tvg-id="Kral Pop TR" tvg-name="Kral Pop TR" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/7v8wmsb/7H8WKei.png" group-title="Music",Kral Pop
https://bqgsd19q.rocketcdn.com/kralpop_720/chunklist.m3u8
#EXTINF:-1 tvg-id="KralTV.tr" tvg-name="Kral TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="" group-title="Music",Kral TV
http://mid5dg6m.rocketcdn.com/kraltv_720/chunklist.m3u8
#EXTINF:-1 tvg-id="Kral TV TR" tvg-name="Kral TV TR" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/j6qp4jT/kral.png" group-title="",Kral TV
https://0ajuu84p.rocketcdn.com/kraltv_720/chunklist.m3u8
#EXTINF:-1 tvg-id="MedMuzik.tr" tvg-name="Med Muzîk" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/t8CpYpP/LcUYuJ5.png" group-title="Music",Med Muzîk
http://137.74.205.201/live3/mm-3mbps.m3u8
#EXTINF:-1 tvg-id="NaturalTV.tr" tvg-name="Natural TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/SBw828K/f460YQg.png" group-title="",Natural TV (1080p)
http://broadcasttr.com:1935/naturaltv/bant1/playlist.m3u8
#EXTINF:-1 tvg-id="NaturalTV.tr" tvg-name="Natural TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/SBw828K/f460YQg.png" group-title="",Natural TV
https://broadcasttr.com:446/naturaltv/bant1/chunklist.m3u8
#EXTINF:-1 tvg-id="NTV TR" tvg-name="NTV TR" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/4Jv1LRS/QP18ELE.png" group-title="News",NTV
https://nt4p9nef.rocketcdn.com/ntvhd.stream_360p/chunklist.m3u8
#EXTINF:-1 tvg-id="NTV TR" tvg-name="NTV TR" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/4Jv1LRS/QP18ELE.png" group-title="News",NTV
https://nt4p9nef.rocketcdn.com/ntvhd.stream_720p/chunklist.m3u8
#EXTINF:-1 tvg-id="ON4TV.tr" tvg-name="ON4 TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/TtbYRZD/i3ua14I.jpg" group-title="",ON4 TV (1080p)
http://yayin.netradyom.com:1935/live/on4/playlist.m3u8
#EXTINF:-1 tvg-id="Rudaw.tr" tvg-name="Rudaw" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/f8Xp0h9/pBgMaFV.jpg" group-title="",Rudaw (1080p)
https://svs.itworkscdn.net/rudawlive/rudawlive.smil/playlist.m3u8
#EXTINF:-1 tvg-id="Star TV HD TR" tvg-name="Star TV HD TR" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/FKfPTrd/star-tv.png" group-title="",Star TV
https://lcgid8xu.rocketcdn.com/startvhd.stream_720p/chunklist.m3u8
#EXTINF:-1 tvg-id="TBMMTV.tr" tvg-name="TBMM TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/vZ2Hw9y/KuDeTzL.png" group-title="",TBMM TV (576p)
https://mecliscanlitv-lh.akamaihd.net/i/MECLISTV_1@127503/master.m3u8
#EXTINF:-1 tvg-id="TGRTBelgesel.tr" tvg-name="TGRT Belgesel" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/VvjSpv8/Tj2G5lV.png" group-title="",TGRT Belgesel (600p)
http://cdn1.streamencoding.com:1935/tgrt_belgesel/amlst:live-30/playlist.m3u8
#EXTINF:-1 tvg-id="tiviTURK.tr" tvg-name="tivi TÜRK" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/9TDs8q8/IFE6LHH.jpg" group-title="",tivi TÜRK
https://stream.tiviturk.de/live/tiviturk.m3u8
#EXTINF:-1 tvg-id="TRT1.tr" tvg-name="TRT 1" tvg-country="TR" tvg-language="Turkish" tvg-logo="" group-title="",TRT 1 (720p)
https://tv-trt1.live.trt.com.tr/master.m3u8
#EXTINF:-1 tvg-id="TRTArabi.tr" tvg-name="TRT Arabi" tvg-country="ARAB" tvg-language="Arabic" tvg-logo="https://i.ibb.co/HTzkq9P/iJIHMq9.png" group-title="News",TRT Arabi (1080p)
https://tv-trtarabi.live.trt.com.tr/master.m3u8
#EXTINF:-1 tvg-id="TRTAvaz.tr" tvg-name="TRT Avaz" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/56kyxmQ/05zMtgQ.png" group-title="",TRT Avaz (720p)
https://tv-trtavaz.live.trt.com.tr/master.m3u8
#EXTINF:-1 tvg-id="TRT Cocuk TR" tvg-name="TRT Cocuk TR" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/c1zC46V/idYZDje.jpg" group-title="Kids",TRT Cocuk (720p)
https://tv-trtcocuk.live.trt.com.tr/master.m3u8
#EXTINF:-1 tvg-id="TRT Haber TR" tvg-name="TRT Haber TR" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/3pKN9gx/BYsl7Nc.jpg" group-title="News",TRT Haber (720p)
https://tv-trthaber.live.trt.com.tr/master.m3u8
#EXTINF:-1 tvg-id="TRTKurdi.tr" tvg-name="TRT Kurdi" tvg-country="TR" tvg-language="Kurdish" tvg-logo="https://i.ibb.co/bB4yG1Y/cWIvunZ.png" group-title="",TRT Kurdi (720p)
https://tv-trtkurdi.live.trt.com.tr/master.m3u8
#EXTINF:-1 tvg-id="TRT Muzik TR" tvg-name="TRT Muzik TR" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/Pmq7bGH/Bi8eMX9.jpg" group-title="Music",TRT Muzik (720p)
https://tv-trtmuzik.live.trt.com.tr/master.m3u8
#EXTINF:-1 tvg-id="TRTTurk.tr" tvg-name="TRT Türk" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/Br1knP7/R8SJH2m.png" group-title="",TRT Türk (720p)
https://tv-trtturk.live.trt.com.tr/master.m3u8
#EXTINF:-1 tvg-id="TRTWorld.tr" tvg-name="TRT World" tvg-country="INT;US;UK;CA;AU;NZ;IE;ZA" tvg-language="English" tvg-logo="" group-title="News",TRT World (1080p)
https://api.trtworld.com/livestream/v1/WcM3Oa2LHD9iUjWDSRUI335NkMWVTUV351H56dqC/master.m3u8
#EXTINF:-1 tvg-id="TRTWorld.tr" tvg-name="TRT World" tvg-country="INT;US;UK;CA;AU;NZ;IE;ZA" tvg-language="English" tvg-logo="https://i.ibb.co/phw4pjP/mjTjJ1N.png" group-title="News",TRT World (1080p)
https://tv-trtworld.live.trt.com.tr/master.m3u8
#EXTINF:-1 tvg-id="UcanKusTV.tr" tvg-name="UçanKuş TV" tvg-country="TR" tvg-language="Turkish" tvg-logo="https://i.ibb.co/qF8Rstv/0PTAYbs.jpg" group-title="",UçanKuş TV
https://ucankus-live.cdnnew.com/ucankus/ucankus.m3u8`)
var segments Segments

func init() {
	segments = Parse(M3uData)
}

func Test_customTags_Get(t *testing.T) {
	channel := segments[0]

	type args struct {
		key string
	}

	tests := []struct {
		name string
		c    customTags
		args args
		want string
	}{
		{name: "tvg-name", c: channel.Ctags, args: args{key: "tvg-name"}, want: "Afyon Turk TV"},
		{name: "tvg-logo", c: channel.Ctags, args: args{key: "tvg-logo"}, want: "https://i.ibb.co/tXbD8Sz/8A1HOMZ.png"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customTags_Set(t *testing.T) {
	channel := segments[0]

	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name    string
		c       customTags
		args    args
		wantErr bool
	}{
		{name: "tvg-country", c: channel.Ctags, args: args{key: "tvg-country", value: "US"}},
		{name: "tvg-name", c: channel.Ctags, args: args{key: "tvg-name", value: "Kanal Adi"}},
		{name: "tvg-language", c: channel.Ctags, args: args{key: "tvg-language", value: "English"}},
		{name: "tvg-logo", c: channel.Ctags, args: args{key: "tvg-logo", value: "http://google.com"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getType(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want MediaType
	}{
		{name: "stream", args: args{"https://5be5d840359c6.streamlock.net/afyonturktv/afyonturktv/playlist.m3u8"}, want: Stream},
		{name: "stream2", args: args{"http://example.com:8080/C@FFFF/keeee/50268"}, want: Stream},
		{name: "stream3", args: args{"http://example.com:8080/user/P090assw99rd/50268.mp4"}, want: Vod},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.args.text); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	f := bytes.NewBufferString(`#EXTM3U
#EXTINF:-1 tvg-id="TV 8,5 HD-TR" group-title="ULUSAL",TR: ATATURK ADD TV
https://5be5d840359c6.streamlock.net/url/afyonturktv/playlist.m3u8`)

	m := Parse(f)

	if m[0].Name != "TR: ATATURK ADD TV" {
		t.Errorf("testParse() = %v, want %v", m[0].Name, "TR: ATATURK ADD TV")
	}

}
