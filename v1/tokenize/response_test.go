package tokenize

import (
	"github.com/pci-bridge/core/pcibtesting"
	"testing"
)

func TestTokenizeWalletResponseValidate(t *testing.T) {
	pcibtesting.UnitTest()

	//r := WalletRequest{}
	//err := json.Unmarshal([]byte(`{"rqid":"01HE0TV6BF1AMWV6QMYRP93B5N","mhid":"39f5bfbd","tsid":"","sessionData":{"colorDepth":24,"javaEnabled":false,"javascriptEnabled":true,"language":"en-GB","screen":{"width":3360,"height":1890},"screenAvailable":{"width":3360,"height":1865},"windowInner":{"width":0,"height":0},"windowOuter":{"width":3360,"height":1865},"timezoneOffsetMins":0,"timezone":"Europe/London","userAgent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36","cookiesEnabled":true,"flashVersion":"","isTouch":false},"bpid":"FID:CST:1689862881:AOYll6q92WDSi","walletType":"GOOGLE_PAY","walletData":"{\"signature\":\"MEQCIC+fxyn0j6WobZ34bzCjW/AI3oXZYV920DalZQORquOCAiBv8w1FBY+mPCRctX44cVsgPDQvJ1aamgrtpYuyJNUIyA\\u003d\\u003d\",\"protocolVersion\":\"ECv1\",\"signedMessage\":\"{\\\"encryptedMessage\\\":\\\"kyoRCtw6pKn8wELnnsinYx6CDKQkDVoR63ykjMXYfMkrHEguiVpTbLsjQlpDpwc0sJNM0dzgtAbQCwwKiowwZOw23Y5U5wBlyrPjEgDXnUPyCDZa3VoqIObHbkGTqRorQWoCMv+rAfSzFGwPcQiKsNfWWaa65MHIQd5ozg/QwyvAynhDucAshGZKpoqdat0Bi5EmzYT62wv/kXhsF1MkwQN+m9M1A84mFv3lPCKQgFiW5zclvzhQZ2UYI99WerpUPxzSIR1CgXya1gRaYDxYmXFeo3BURShjbYE+gi0lXM9tDCsVI79C16lVtL4KasSBAuknDJaaPu7/lIIJJpguFt26xTqOYpLemZovMBgWz3fLtNF2ukZcAzKluaUVrciheFO7/EdKmSLJpl8Kaa6ryi6Qq9k+LsjplxMWl/ioefm/OSo3bkrICSlIgShqo0LEiZYjWCtKqoB429L3BiZUBx1NEQAxuGFGnw1vlOP/\\\",\\\"ephemeralPublicKey\\\":\\\"BKcW+SXCMgwDZbU6WNwfor0Y7WqU5j26B3V45LoSGfvqyAK1bPFXA7IypxmBy4kRBATBvs+rkARWk2evggQuJE0\\\\u003d\\\",\\\"tag\\\":\\\"Q/CzmEm22g9Qo4vnbvPiw+QqMx9HoyUDcqdiU5VuCN0\\\\u003d\\\"}\"}"}`), &r)
	//assert.Nil(t, err)
}
