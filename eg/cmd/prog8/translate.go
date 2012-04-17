package main

import "crypto/sha512"
import "fmt"
import "image"
import "image/png"

import "github.com/mewmew/playground/pic"

/// ### todo ###
///   - Only single characters are present in the map.
///   - The lookup will fail for character sequences such as "AW":
///     (09A32E1C2EE7E7BBBF36049601B02B7EA4B48C64CCB487F9708999015771799470486B0C02805B2FC9DD9563D7A2E041E44304FACED37233B20E2077380B58DE)
/// ############
// table is a map from image content hash sum to a letter or number.
var table = map[string]string{
   "F73E6481F7D74DBD9361EA482954FBCBE1762B7A40C650E21A4C963919323716F31A72D61F01C7B0095643CCCF98BA4371CD8DA8D4B9B5518696436465557ED3": "0",
   "CD1E456C3990722E02D839790D68A34B75C3B5B2E84118E609F4ABBBC4A05891662CB4C5642A35E073EFD935C551B88CDFCCC00A44B9D95B1539992409ADD13E": "1",
   "B37D6348137E7006C76377A60A0D7B4CC5F8164358B29084F95613E516CD580705587108DA504D701E5B2D659C5466F5FC331720E652761E07108A0FB43FE489": "2",
   "A0AF46198BEE9A8E26457EC75C51A82E5053284AE7266BF998C7A10435F0EB522BC9BCEC4DF064068E7A61F4FFBC54E75367356ED4549811A410C1F451537555": "3",
   "F66445EE1A502E2E0AC858D37F02C7719CF6EFC81994ABCD56FAAAF0A4658C0E1A91EDDDA0501D1426EE5106EBF862D05615AD2BED6F73B05EF2FB0807C4956E": "4",
   "4FB551B394F9F5AEA9750FD5FAF3B78E3A938B3A3E13994F4DC6A35F1E16383F136D4CFC8583FEC72D9AAC708456A681568DDCA2A03AB549682D3B5CFDB43273": "5",
   "294346D9FB222B94493EAD328E94498BFCAE041E464FA538AC4C8ED8E6273D701BF79B1E1E5EBD68140DF327BF75DA7F30F93CC640814AF58496FB4472108D8B": "6",
   "8F21E3EC2577218D395A257191F4E7D26773E93E9F7AF9047AB886D7EEB159D32991C3A09D4243E5643657BB3A28E3A95CABC0ED907E612050C506886A1692DE": "7",
   "2839CEC5B7CFA816B78830BE15E0271500A0E7EF7ACC0010D47DB68B0839159B1705F3059139A74E1C21B69620959557E2A0C7EC1CC9D1A88A387EB23792ED49": "8",
   "DB7B6454FA9D7EF4D7AD0109EBE55519F8A124B2A7AAFA43032933DE2769B616CFFA6F31C66E362A17E6A7367D6DC5318EA80726E24027355184EB3EA2B504EB": "9",
   "EA26B7ADA53A34915611D72AB223EC55D2D3C1856585F9ABA16E84C76D9F30FBC98D84012FFD1540A711F2ABDD8E753C5E97E401C8757BFA0D3051A392100E3A": "A",
   "CDE886D0E887FA5A4EEFE45E9209257204944CCBC74F7592201D23589BFE84901A60107897C660310F635A991234157BB3CEE64D6E70645077B2A673212A0DF1": "B",
   "C055ADB5B31124E959365724A56A320CBB4500EBC5530707D6E01D706E08C379EDBB97ADC62BF3DE56AA4A4D6392EA85B11789776248953BF873BA1C2BD33F4B": "C",
   "658094110E5327E033623897B9E77C8D83E6F6443761DEB2F8BD02982A3BC95BAA017BABA90240508F79C3B1E86AA306A0261C8387BBD2EFB1297A6632DF0C32": "D",
   "D847A458748EF44DE2E14CBCECBCA26B817629B004BEA3E224CF4722316DB99977928948F215D4AE2E73033D98DD9CA2C639649797A0A57E3E1039464C42D1D3": "E",
   "4131D76B431E1015D52342266E37ED31DC7F20910F5C46507611BF7C50917D6E62CC6982541257A27D757C9FCC859A453447E97728B86694F054466FA3CB4399": "F",
   "289A020420E0391FD59F2076858AF74C6BA8843BD7A19F6A90886367DB6572ADF91FA25399B0E40A5F1B648ED29353182A5B661E97FBBD1D88C0116D326EBBC0": "G",
   "FE0F36DF27F8F183CC10DDAE8B0311B66D9D6790E536BBE431ED25011BFD9468C246B9E3CB2B808845A94B717AD7EA7F12F8853F6B6DA723F59691F7229BB635": "H",
   "59E5ADE87BE2F0D8BA12C14FB470221516CF3EB8D77EB1796B07EEAC30B5C486D86261B8680D4CE7DA862D104B8B036EB2091E244AE97BC022916EA69002690A": "I",
   "A82E30F090173D471E83122490BD25BBDF94341DBD08C04A6B23D46ECFD56F2910ACDFF9EA997874111F914F8A0AF3778C34707C105870430AC2872F295800B3": "J",
   "42FD9E238ED3FA93801085D699449F7C92145564CEE6C9C32D18CA0AECDDF8FBF50BAC63DE239FEA444D79B8B23EA7E3F11EE6B0154988008E37DEE748654BEB": "K",
   "9284A3C4B23FB17A9013026A603BB83CE31D9983181BBD9045525EF00A45CD7B56FE7721C9700DF82AA6F7490B13E6AA26DBDAAA9E7D788EDB2A87851B0F7B0F": "L",
   "846CD7C98981B533B0044320D6CC7F58C521312FBCD5455581217BFA4E70D0E3F5AA4917A44959D677DF182B4E8B610F36B93102A5A0BC8BB0560A685111A032": "M",
   "6C4E1BB85201CC22EAD298C812AADF04B5ED80EC4AD79ABBC8448FCBC854D83023B63F58462037AC92D7EACAF78A0CA34CBE40F08FA4A27B54535C877F5FB20C": "N",
   "2B31F63901FA95F60C28FC11F941FF55D603896E667733E02AEC08AD5F7E7BA4BDE463334DF0B54D35C1A7B2C20DB2EB94A0194216DE2536E8E9B333038E08D9": "O",
   "4CAA3235BE3F05ABC491309FBEB7D02CED010E3E2E8CFE15C8A9267E6622F18ED535DEAD39C1EC072D03D734641B6F7821B70C8E3B97C21787813E4E998B9BC1": "P",
   "5E4F12A5C1263341C41BD2473513840458CD22FBDD3F08B34852204AB74DDBB32D880EA4ED39F2DF3F738CF9DB08BC3A4F763E5ABC0103F795D044A785DEF5BF": "Q",
   "9124AE9842BEABBFFC6C5E17A56962CADD6FA1A2AE20556AAA1A08893EE9EC7AF06E0A3707D8EBCE6F59978753F167431226433A50DCDB1A3B58143101C1D5E5": "R",
   "8DF7D0422DB18555C29B62DECDEF7C22D9A6C84516032978B45B16D72A9D6CAF1E77AB5DA0C6B315704CE109EBEE4169039DAC9EE70CCB777D9925D962184569": "S",
   "3F606F38BEAD5397CCF75E406855673EA704E705A2D8AC42BC373347E13740C3D381BB4AF60F52899628CC90A0A23E0BD0DE5B6D3CFB45E5A41F8F55A43AF5AE": "T",
   "ED7400B850B21AA3483F8490A4110F67588DD6C807AE05DD3A5DC99EEC08B4D2B69691DAF4D8FACD337151F26EA3F0D63D326C05935D1BB482818D02EAF8336A": "U",
   "4F3FE263BF88AF3DD52B14BF6F2F758753143617BFC200651BE2D37C78B813D684E186700324146BFFC71DDB435445CBD8DF659F7CBBDF8BA4E0884195F040A1": "V",
   "C754CBCC6062B5F71249A8D53ABB8F9AE66F1D1980B3F734C6891441403B30E933F8D4DCEB1448E5B2B555E0DC9EB3D68436A21A5D76D290BB3318320C8ABECF": "W",
   "F8B3FA2F0DAD2BD9C797A0A6955BB20F89BE7DDB6E36F5B88A04E9813B9A7F4327B8717FA3E68688A0252FE4BBFE174E2D75E95DBCF54B44AFD0E0AB59D6B6EA": "X",
   "B44BEFD81268CDFEF2AC03793BB7D7C2A3FC4566154AAEEABB4DBB2A9D72D00F0B2EF44365146C59E8E45346E0F46F47628DCFDB0DCF9C232FBF0A0F7BA5A22C": "Y",
   "A9EF04F8CA6FBED179EF4396855E40ABA720AE062DC59D6FE493375B022041085430B8EC6C7DCDAC277C82B822F2A159D940B4836132B1A5EF7546B32DC6762F": "Z",
}

// translate translates each sub image back into a character, by compare them
// against a stored version, and returns a string composed by these characters.
func translate(subs []pic.SubImager) (text string, err error) {
   for _, sub := range subs {
      hashSum, err := getHashSum(sub)
      if err != nil {
         return "", err
      }
      c, ok := table[hashSum]
      if !ok {
         return "", fmt.Errorf("table doesn't contain the content hash sum ('%s').", hashSum)
      }
      text += c
   }
   return text, nil
}

// getHashSum returns a hash sum based on the image's content.
func getHashSum(img image.Image) (hashSum string, err error) {
   hash := sha512.New()
   err = png.Encode(hash, img)
   if err != nil {
      return "", err
   }
   hashSum = fmt.Sprintf("%X", hash.Sum(nil))
   return hashSum, nil
}
