/**********************************************************************
 * Unicode Test                                                       *
 * Tomi Häsä                                                          *
 * tomi.hasa@gmail.com                                                *
 * v1.00 - 2017-01-09                                                 *
 *********************************************************************/

/**********************************************************************
 * Unicode test for Google Chrome, Google Plus, Android, iPhone, etc. *
 *********************************************************************/


package main

import (

    "fmt"
)


const (

    // version information
    UNICODETEST  string = "Unicode Test v1.00 - 2017-01-09"

    // Unicode version information is not totally accurate as it takes a long time to check
    U0  int = 0   // Unicode version not resolved
    U1  int = 1   // Unicode 1
    U2  int = 2   // Unicode 2
    U3  int = 3   // Unicode 3
    U4  int = 4   // Unicode 4
    U5  int = 5   // Unicode 5
    U6  int = 6   // Unicode 6
    U7  int = 7   // Unicode 7
    U8  int = 8   // Unicode 8 or less (estimate)
    U9  int = 9   // Unicode 9 or more (estimate)

    // title for a group
    EMPTY       string = "-"      // not a title, because not the first area in a group
    SPACES      string = "  "     // two spaces between characters as one space is not always enough
    COMMASPACE  string = ", "     // comma after the last ucode of the area in a title
    NEWLINES    string = "\n\n"   // newlines after the last ucode of the last area of the group in a title
    HYPHEN      string = "-"      // hyphen between ucodes in a title
    OPENINGPARENTHESIS  string = "("   // before the first ucode in the group of areas in a title
    CLOSINGPARENTHESIS  string = ")"   // after the last ucode of the group of areas in a title

    // areas in a group
    MAX       int = 20   // maximum number of areas in a group
    SOLO      int = 0    // the one and only area in a group
    TOP       int = 1    // first area in a group
    CONTINUE  int = 2    // group of areas continues
    END       int = 3    // group of areas ends
)


type Area struct {        // Unicode area

    start        int      // character in hex (base-16, 0...f)
    end          int      // character in hex (base-16, 0...f)
    version      int      // Unicode version (usually an estimate)
    continues    int      // group continues in the next Unicode area
    description  string   // description of the Unicode area
}


var (

    index1       int    = 0
    value1       Area
    index2       int    = 0
    value2       Area
    index3       int    = 0
    value3       Area
    description  string = ""
    start        string = ""
    end          string = ""
    onearea      Area   = Area{}
    group        []Area = make( []Area, 0, MAX )
)


func main() {

    // these lines have not been split to increase readability
    groups := [...]Area{

        { '\u0021', '\u007E', U1, SOLO, "Basic Latin" },         // !...~

        { '\u00C0', '\u00F6', U1, TOP, "Latin 1 Supplement" },   // À...ö
        { '\u00F8', '\u00FF', U1, END, EMPTY },                  // ø...ÿ

        { '\u0100', '\u017F', U1, SOLO, "Latin Extended A" },    //  ...ſ

        { '\u0180', '\u01F5', U1, TOP,      "Latin Extended B" },   // ƀ...ǵ
        { '\u01F6', '\u01F9', U2, CONTINUE, EMPTY },                // Ƕ...ǹ
        { '\u01FA', '\u0217', U1, CONTINUE, EMPTY },                // Ǻ...ȗ
        { '\u0218', '\u021F', U2, CONTINUE, EMPTY },                // Ș...ȟ
        { '\u0220', '\u0221', U1, CONTINUE, EMPTY },                // Ƞ...ȡ
        { '\u0222', '\u0233', U2, CONTINUE, EMPTY },                // Ȣ...ȳ
        { '\u0234', '\u024F', U1, END,      EMPTY },                // ȴ...ɏ

        { '\u1E00', '\u1EFF', U1, SOLO, "Latin Extended Additional" },   //  ...ỿ

        { '\u249C', '\u24E9', U1, SOLO, "Enclosed Alphabetics" },        // ⒜...ⓩ


        { '\u0370', '\u0377', U1, TOP,      "Greek" },   // Ͱ...ͷ
        { '\u037A', '\u037F', U1, CONTINUE, EMPTY },     // ͺ...Ϳ
        { '\u0384', '\u038A', U1, CONTINUE, EMPTY },     // ΄...Ί
        { '\u038C', '\u038C', U1, CONTINUE, EMPTY },     // Ό
        { '\u038E', '\u03A1', U1, CONTINUE, EMPTY },     // Ύ...Ρ
        { '\u03A3', '\u03E1', U1, END,      EMPTY },     // Σ...ϡ


        { '\u0410', '\u044F', U1, TOP,      "Cyrillic" },   // А...я
        { '\u0400', '\u040F', U1, CONTINUE, EMPTY },        //  ...Џ
        { '\u0450', '\u045F', U1, CONTINUE, EMPTY },        // ѐ...џ
        { '\u048A', '\u04F9', U1, END,      EMPTY },        // Ҋ...ӹ


        { '\u05D0', '\u05EA', U1, TOP, "Hebrew" },   // א...ת
        { '\u05F0', '\u05F4', U1, END, EMPTY },      // װ...״


        { '\u0621', '\u064A', U1, TOP,      "Arabic" },   // ...
        { '\u0660', '\u066C', U1, CONTINUE, EMPTY },      // ...
        { '\u0671', '\u0672', U1, CONTINUE, EMPTY },      // ...
        { '\u0674', '\u06D3', U1, END,      EMPTY },      // ...


        { '\u0250', '\u02AF', U1, SOLO, "IPA" },      // ɐ...ʯ

        { '\u1D00', '\u1D7F', U1, TOP, "Non-IPA" },   //  ...ᵿ
        { '\u1D80', '\u1DBF', U1, END, EMPTY },       // ᶀ...ᶿ


        { '\u0021', '\u002F', U1, TOP,      "Punctuation: Basic Latin" },   // !.../
        { '\u003A', '\u003F', U1, CONTINUE, EMPTY },                        // :...?
        { '\u0060', '\u0060', U1, CONTINUE, EMPTY },                        // `
        { '\u007B', '\u007E', U1, END,      EMPTY },                        // {...~

        { '\u00A1', '\u00A1', U1, TOP,      "Punctuation: Latin 1 Supplement" },   // ¡
        { '\u00A6', '\u00A8', U1, CONTINUE, EMPTY },                               // ¦...¨
        { '\u00AA', '\u00AB', U1, CONTINUE, EMPTY },                               // ª...«
        { '\u00AF', '\u00B0', U1, CONTINUE, EMPTY },                               // ¯...°
        { '\u00B2', '\u00B4', U1, CONTINUE, EMPTY },                               // ²...´
        { '\u00B6', '\u00BB', U1, CONTINUE, EMPTY },                               // ¶...»
        { '\u00BF', '\u00BF', U1, END,      EMPTY },                               // ¿

        { '\u2010', '\u2010', U1, TOP,      "General Punctuation" },   // ‐
        { '\u2012', '\u2027', U1, CONTINUE, EMPTY },                   // ‒...‧
        { '\u2030', '\u205E', U1, END,      EMPTY },                   // ‰...⁞

        { '\u2E00', '\u2E42', U8, TOP, "Supplemental Punctuation" },   // ⸀...⹂
        { '\u2E43', '\u2E44', U9, END, EMPTY },                        // ⹃...⹄


        { '\u0030', '\u0039', U1, SOLO, "Numbers" },              // 0...9

        { '\u00BC', '\u00BE', U1, SOLO, "Numbers: Fractions" },   // ¼...¾

        { '\u2150', '\u218B', U1, SOLO, "Number Forms" },         // ⅐...↋

        { '\u2460', '\u249B', U1, TOP, "Enclosed Numbers" },      // ①...⒛
        { '\u24EA', '\u24FF', U1, END, EMPTY },                   // ⓪...⓿

        { '\u2776', '\u2793', U1, SOLO, "Numbers: Dingbats" },    // ❶...➓


        { '\u00B2', '\u00B3', U1, TOP,      "Superscripts" },   // ²...³
        { '\u00B9', '\u00B9', U1, CONTINUE, EMPTY },            // ¹
        { '\u2070', '\u2071', U1, CONTINUE, EMPTY },            // ⁰...i
        { '\u2074', '\u207F', U1, END,      EMPTY },            // ⁴...ⁿ

        { '\u2080', '\u208E', U1, TOP, "Subscripts" },          // ⁰...₎
        { '\u2090', '\u209C', U1, END, EMPTY },                 // ₐ...ₜ


        { '\u0021', '\u0021', U1, TOP,      "Mathematics" },   // !
        { '\u0025', '\u0025', U1, CONTINUE, EMPTY },           // %
        { '\u0028', '\u002F', U1, CONTINUE, EMPTY },           // (.../
        { '\u003A', '\u003E', U1, CONTINUE, EMPTY },           // :...>
        { '\u005B', '\u005E', U1, CONTINUE, EMPTY },           // [...^
        { '\u007B', '\u007E', U1, CONTINUE, EMPTY },           // {...~
        { '\u00AC', '\u00AC', U1, CONTINUE, EMPTY },           // ¬
        { '\u00B0', '\u00B1', U1, CONTINUE, EMPTY },           // °...±
        { '\u00B4', '\u00B4', U1, CONTINUE, EMPTY },           // ´
        { '\u00B7', '\u00B7', U1, CONTINUE, EMPTY },           // ·
        { '\u00D7', '\u00D7', U1, END,      EMPTY },           // ×

        { '\u27C0', '\u27EF', U8, SOLO, "Mathematical Symbols A" },   // ⟀...⟯

        { '\u2980', '\u29FF', U8, SOLO, "Mathematical Symbols B" },   // ⦀...⧿

        { '\u2A00', '\u2AFF', U8, SOLO, "Mathematics: Operators" },   // ⨀...⫿


        { '\u2103', '\u2103', U0, TOP,      "Natural Sciences" },   // ℃
        { '\u2107', '\u2107', U0, CONTINUE, EMPTY },                // ℇ
        { '\u2109', '\u2115', U0, CONTINUE, EMPTY },                // ℉...ℕ
        { '\u2118', '\u211D', U0, CONTINUE, EMPTY },                // ℘...ℝ
        { '\u2124', '\u2134', U0, END,      EMPTY },                // ℤ...ℴ


        { '\u2300',     '\u23FA',     U8, TOP, "Technical" },        // ⌀...⏺
        { '\u23FB',     '\u23FE',     U9, END, EMPTY },              // ⏻...⏾

        { '\U0001F19B', '\U0001F1AC', U9, SOLO, "Technical: TV" },   // 🆛...🆬


        { '\u0024', '\u0024', U1, TOP,      "Currency" },   // $
        { '\u00A2', '\u00A5', U1, CONTINUE, EMPTY },        // ¢...¥
        { '\u20A0', '\u20AB', U8, CONTINUE, EMPTY },        // ₠...₫
        { '\u20AC', '\u20AC', U2, CONTINUE, EMPTY },        // €
        { '\u20AD', '\u20BE', U8, END,      EMPTY },        // ₭...₾


        { '\u2100', '\u2101', U1, TOP,      "Business" },   // ℀...℁
        { '\u2105', '\u2106', U1, CONTINUE, EMPTY },        // ℅...℆
        { '\u2120', '\u2121', U1, CONTINUE, EMPTY },        // ℠...℡
        { '\u213B', '\u213B', U1, END,      EMPTY },        // ℻


        { '\u00A9',     '\u00A9',     U1, TOP,      "Copyright" },   // ©
        { '\u00AE',     '\u00AE',     U1, CONTINUE, EMPTY },         // ®
        { '\u2117',     '\u2117',     U1, CONTINUE, EMPTY },         // ℗
        { '\u2122',     '\u2122',     U1, CONTINUE, EMPTY },         // ™
        { '\U0001F12E', '\U0001F12E', U8, END,      EMPTY },         // 🄮


        { '\u211E', '\u211E', U1, SOLO, "Health" },   // ℞


        { '\u0023', '\u0023', U1, TOP, "Social Media" },   // #
        { '\u0040', '\u0040', U1, END, EMPTY },            // @


        { '\u2600', '\u26FF', U8, SOLO, "Symbols" },       // ☀...⛿

        { '\u2B00', '\u2B73', U8, TOP,      "Symbols" },   // ⬀...⭳
        { '\u2B76', '\u2B95', U8, CONTINUE, EMPTY },       // ⭶...⮕
        { '\u2B98', '\u2BB9', U8, CONTINUE, EMPTY },       // ⮘...⮹
        { '\u2BBD', '\u2BC8', U8, CONTINUE, EMPTY },       // ⮽...⯈
        { '\u2BCA', '\u2BD1', U8, CONTINUE, EMPTY },       // ⯊...⯑
        { '\u2BEC', '\u2BEF', U8, END,      EMPTY },       // ⯬...⯯

        { '\U0001F300', '\U0001F579', U8, TOP,      "Symbols and Pictographs" },   // 🌀...🕹
        { '\U0001F57A', '\U0001F57A', U9, CONTINUE, EMPTY },                       // 🕺
        { '\U0001F57B', '\U0001F5A3', U8, CONTINUE, EMPTY },                       // 🕻...🖣
        { '\U0001F5A4', '\U0001F5A4', U9, CONTINUE, EMPTY },                       // 🖤
        { '\U0001F5A5', '\U0001F5FF', U8, END,      EMPTY },                       // 🖥...🗿

        { '\U0001F910', '\U0001F918', U8, TOP,      "Supplemental Symbols and Pictographs" },   // 🤐...🤘
        { '\U0001F919', '\U0001F91E', U9, CONTINUE, EMPTY },   // 🤙...
        { '\U0001F9C0', '\U0001F9C0', U8, CONTINUE, EMPTY },   // 🧀
        { '\U0001F920', '\U0001F927', U9, CONTINUE, EMPTY },   // 🤠...🤧
        { '\U0001F930', '\U0001F930', U9, CONTINUE, EMPTY },   // 🤰
        { '\U0001F933', '\U0001F93E', U9, CONTINUE, EMPTY },   // 🤳...🤾
        { '\U0001F940', '\U0001F94B', U9, CONTINUE, EMPTY },   // 🥀...🥋
        { '\U0001F950', '\U0001F95E', U9, CONTINUE, EMPTY },   // 🥐...🥞
        { '\U0001F980', '\U0001F984', U8, CONTINUE, EMPTY },   // 🦀...🦄
        { '\U0001F985', '\U0001F991', U9, END,      EMPTY },   // 🦅...🦑


        { '\u2700',     '\u27BF',     U1, SOLO, "Dingbats" },              // ✀...➿

        { '\U0001F650', '\U0001F67F', U8, SOLO, "Ornamental Dingbats" },   // 🙐...🙿


        { '\u2190', '\u21FF', U1, SOLO, "Arrows" },                  // ←...⇿

        { '\u27F0', '\u27FF', U1, SOLO, "Supplemental Arrows A" },   // ⟰...⟿

        { '\u2900', '\u297F', U1, SOLO, "Supplemental Arrows B" },   // ⤀...⥿

        { '\u2794', '\u2794', U1, TOP,      "Arrows: Dingbats" },    // ➔
        { '\u2798', '\u27AF', U1, CONTINUE, EMPTY },                 // ➘...➯
        { '\u27B1', '\u27BE', U1, END,      EMPTY },                 // ➱...➾


        { '\u2B00', '\u2B11', U1, TOP,      "Arrows" },   // ⬀...⬑
        { '\u2B30', '\u2B4F', U1, CONTINUE, EMPTY },      // ⬰...⭏
        { '\u2B5A', '\u2B73', U1, CONTINUE, EMPTY },      // ⭚...⭳
        { '\u2B76', '\u2B95', U1, CONTINUE, EMPTY },      // ⭶...⮕
        { '\u2B98', '\u2BB9', U1, CONTINUE, EMPTY },      // ⮘...⮹
        { '\u2BC5', '\u2BC8', U1, CONTINUE, EMPTY },      // ⯅...⯈
        { '\u2BEC', '\u2BEF', U1, END,      EMPTY },      // ⯬...⯯


        { '\u25A0', '\u25FF', U1, SOLO, "Geometric Shapes" },   // ■...◿

        { '\U0001F532', '\U0001F539', U0, SOLO, "Geometric Shapes in Symbols and Pictographs" },   // 🔲...🔹


        { '\u2500', '\u257F', U1, SOLO, "Box Drawing" },   // ─...╿


        { '\u2580', '\u259F', U1, SOLO, "Block Elements" },   // ▀...▟


        { '\u2639',     '\u263B',     U1, SOLO, "\"Emoticons\"" },   // ☹...☻

        { '\U0001F600', '\U0001F640', U6, TOP,      "Emoticons" },   // 😀...🙀
        { '\U0001F641', '\U0001F642', U7, CONTINUE, EMPTY },         // 🙁...🙂
        { '\U0001F643', '\U0001F644', U8, CONTINUE, EMPTY },         // 🙃...🙄
        { '\U0001F645', '\U0001F64F', U6, END,      EMPTY },         // 🙅...🙏


        { '\U0001F466', '\U0001F483', U0, TOP,      "Emoticons in Symbols and Pictographs" },   // 👦...💃
        { '\U0001F491', '\U0001F491', U0, CONTINUE, EMPTY },   // 💑
        { '\U0001F4A1', '\U0001F4AD', U0, CONTINUE, EMPTY },   // 💡...💭
        { '\U0001F451', '\U0001F463', U0, END,      EMPTY },   // 👑...👣


        { '\U00010080', '\U000100FA', U8, SOLO, "Ideograms" },   // 𐂀...𐃺


        { '\U00013000', '\U0001342E', U0, SOLO, "Hieroglyphs: Egyptian" },   // 𓀀...𓐮

    }

    fmt.Printf( "%v%v", UNICODETEST, NEWLINES )

    for index1, value1 = range groups {

        onearea = value1

        // add a new area to a group
        group = append( group, onearea )

        if ( value1.continues == SOLO || value1.continues == TOP ) {

            titledescription( value1 )
            fmt.Printf( OPENINGPARENTHESIS )
        }

        if ( value1.continues == TOP || value1.continues == CONTINUE ) {

            titlecodes( value1 )
            fmt.Printf( COMMASPACE )
        }

        if ( value1.continues == SOLO || value1.continues == END ) {

            titlecodes( value1 )
            fmt.Printf( "%v%v", CLOSINGPARENTHESIS, NEWLINES )

            for index2, value2 = range group {

                charactersforarea( value2 )
            }

            // newlines after the last character
            fmt.Printf( NEWLINES )

            // reset after the last area has been added and printed
            // length of the new slice will be zero and capacity MAX
            group = make( []Area, 0, MAX )

            // Do not reset this way:
            // group = nil
        }
    }

    // when the program ends
    fmt.Printf( NEWLINES )
}


func titledescription( area Area ) {

    description := area.description

    fmt.Printf( "%v ", description )

}


func codefortitle( character int ) {

    fmt.Printf( "U" )

    // not more than 0xff i.e. less than 0x100
    if ( character < 255 ) {

        fmt.Printf( "00" )
    }

    fmt.Printf( "%X", character )
}


func titlecodes( area Area ) {

    start  := area.start
    end    := area.end

    if ( start == end ) {

       codefortitle( start )

    } else {

        codefortitle( start )
        fmt.Printf( HYPHEN )
        codefortitle( end )
    }
}


func charactersforarea( area Area ) {

    start  := area.start
    end    := area.end
    //version := area.version

    //fmt.Printf( "U%d: ", version )

    for i := start; i < (end + 1); i++ {

        fmt.Printf( "%c%v", i, SPACES )
    }
}


func testprintgroup() {

    length    := len( group )
    capacity  := cap( group )

    fmt.Printf( "\nlength   = %d ", length )
    fmt.Printf( "\ncapacity = %d ", capacity )

    for index3, value3 = range group {
    
        fmt.Printf( "%c ", index3 )
    }

    fmt.Printf( NEWLINES )
}
