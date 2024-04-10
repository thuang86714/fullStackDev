// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__f32toa = 31168
    _entry__f64toa = 192
    _entry__format_significand = 35344
    _entry__format_integer = 3040
    _entry__get_by_path = 25696
    _entry__fsm_exec = 17920
    _entry__advance_string = 14320
    _entry__advance_string_default = 36736
    _entry__do_skip_number = 20528
    _entry__skip_one_fast = 22176
    _entry__html_escape = 8912
    _entry__i64toa = 3472
    _entry__u64toa = 3584
    _entry__lspace = 16
    _entry__quote = 4864
    _entry__skip_array = 17872
    _entry__skip_number = 21792
    _entry__skip_object = 20160
    _entry__skip_one = 21952
    _entry__unquote = 6576
    _entry__validate_one = 22000
    _entry__validate_utf8 = 29904
    _entry__validate_utf8_fast = 30576
    _entry__value = 12320
    _entry__vnumber = 15648
    _entry__atof_eisel_lemire64 = 10160
    _entry__atof_native = 11712
    _entry__decimal_to_f64 = 10528
    _entry__right_shift = 36304
    _entry__left_shift = 35808
    _entry__vsigned = 17200
    _entry__vstring = 14144
    _entry__vunsigned = 17520
)

const (
    _stack__f32toa = 48
    _stack__f64toa = 80
    _stack__format_significand = 24
    _stack__format_integer = 16
    _stack__get_by_path = 280
    _stack__fsm_exec = 168
    _stack__advance_string = 64
    _stack__advance_string_default = 64
    _stack__do_skip_number = 48
    _stack__skip_one_fast = 176
    _stack__html_escape = 72
    _stack__i64toa = 16
    _stack__u64toa = 8
    _stack__lspace = 8
    _stack__quote = 56
    _stack__skip_array = 176
    _stack__skip_number = 104
    _stack__skip_object = 176
    _stack__skip_one = 176
    _stack__unquote = 88
    _stack__validate_one = 176
    _stack__validate_utf8 = 48
    _stack__validate_utf8_fast = 24
    _stack__value = 328
    _stack__vnumber = 240
    _stack__atof_eisel_lemire64 = 32
    _stack__atof_native = 136
    _stack__decimal_to_f64 = 80
    _stack__right_shift = 8
    _stack__left_shift = 24
    _stack__vsigned = 16
    _stack__vstring = 120
    _stack__vunsigned = 8
)

const (
    _size__f32toa = 3392
    _size__f64toa = 2848
    _size__format_significand = 464
    _size__format_integer = 432
    _size__get_by_path = 4208
    _size__fsm_exec = 1692
    _size__advance_string = 1280
    _size__advance_string_default = 944
    _size__do_skip_number = 924
    _size__skip_one_fast = 3016
    _size__html_escape = 1248
    _size__i64toa = 48
    _size__u64toa = 1232
    _size__lspace = 144
    _size__quote = 1696
    _size__skip_array = 48
    _size__skip_number = 160
    _size__skip_object = 48
    _size__skip_one = 48
    _size__unquote = 2272
    _size__validate_one = 48
    _size__validate_utf8 = 672
    _size__validate_utf8_fast = 560
    _size__value = 1308
    _size__vnumber = 1552
    _size__atof_eisel_lemire64 = 368
    _size__atof_native = 608
    _size__decimal_to_f64 = 1184
    _size__right_shift = 400
    _size__left_shift = 496
    _size__vsigned = 320
    _size__vstring = 128
    _size__vunsigned = 336
)

var (
    _pcsp__f32toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {3350, 48},
        {3351, 40},
        {3353, 32},
        {3355, 24},
        {3357, 16},
        {3359, 8},
        {3363, 0},
        {3385, 48},
    }
    _pcsp__f64toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {2788, 56},
        {2792, 48},
        {2793, 40},
        {2795, 32},
        {2797, 24},
        {2799, 16},
        {2801, 8},
        {2805, 0},
        {2843, 56},
    }
    _pcsp__format_significand = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {452, 24},
        {453, 16},
        {455, 8},
        {457, 0},
    }
    _pcsp__format_integer = [][2]uint32{
        {1, 0},
        {4, 8},
        {412, 16},
        {413, 8},
        {414, 0},
        {423, 16},
        {424, 8},
        {426, 0},
    }
    _pcsp__get_by_path = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {4012, 104},
        {4016, 48},
        {4017, 40},
        {4019, 32},
        {4021, 24},
        {4023, 16},
        {4025, 8},
        {4026, 0},
        {4194, 104},
    }
    _pcsp__fsm_exec = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1342, 104},
        {1346, 48},
        {1347, 40},
        {1349, 32},
        {1351, 24},
        {1353, 16},
        {1355, 8},
        {1356, 0},
        {1692, 104},
    }
    _pcsp__advance_string = [][2]uint32{
        {14, 0},
        {18, 8},
        {20, 16},
        {22, 24},
        {24, 32},
        {26, 40},
        {27, 48},
        {557, 56},
        {561, 48},
        {562, 40},
        {564, 32},
        {566, 24},
        {568, 16},
        {570, 8},
        {571, 0},
        {1268, 56},
    }
    _pcsp__advance_string_default = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {552, 64},
        {556, 48},
        {557, 40},
        {559, 32},
        {561, 24},
        {563, 16},
        {565, 8},
        {566, 0},
        {931, 64},
    }
    _pcsp__do_skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {849, 48},
        {850, 40},
        {852, 32},
        {854, 24},
        {856, 16},
        {858, 8},
        {859, 0},
        {924, 48},
    }
    _pcsp__skip_one_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {658, 176},
        {659, 168},
        {661, 160},
        {663, 152},
        {665, 144},
        {667, 136},
        {671, 128},
        {3016, 176},
    }
    _pcsp__html_escape = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1224, 72},
        {1228, 48},
        {1229, 40},
        {1231, 32},
        {1233, 24},
        {1235, 16},
        {1237, 8},
        {1239, 0},
    }
    _pcsp__i64toa = [][2]uint32{
        {14, 0},
        {34, 8},
        {36, 0},
    }
    _pcsp__u64toa = [][2]uint32{
        {1, 0},
        {161, 8},
        {162, 0},
        {457, 8},
        {458, 0},
        {756, 8},
        {757, 0},
        {1221, 8},
        {1223, 0},
    }
    _pcsp__lspace = [][2]uint32{
        {1, 0},
        {89, 8},
        {90, 0},
        {103, 8},
        {104, 0},
        {111, 8},
        {113, 0},
    }
    _pcsp__quote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1649, 56},
        {1653, 48},
        {1654, 40},
        {1656, 32},
        {1658, 24},
        {1660, 16},
        {1662, 8},
        {1663, 0},
        {1690, 56},
    }
    _pcsp__skip_array = [][2]uint32{
        {1, 0},
        {28, 8},
        {34, 0},
    }
    _pcsp__skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {107, 56},
        {111, 48},
        {112, 40},
        {114, 32},
        {116, 24},
        {118, 16},
        {120, 8},
        {121, 0},
        {145, 56},
    }
    _pcsp__skip_object = [][2]uint32{
        {1, 0},
        {28, 8},
        {34, 0},
    }
    _pcsp__skip_one = [][2]uint32{
        {1, 0},
        {30, 8},
        {36, 0},
    }
    _pcsp__unquote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1684, 88},
        {1688, 48},
        {1689, 40},
        {1691, 32},
        {1693, 24},
        {1695, 16},
        {1697, 8},
        {1698, 0},
        {2270, 88},
    }
    _pcsp__validate_one = [][2]uint32{
        {1, 0},
        {35, 8},
        {41, 0},
    }
    _pcsp__validate_utf8 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {623, 48},
        {627, 40},
        {628, 32},
        {630, 24},
        {632, 16},
        {634, 8},
        {635, 0},
        {666, 48},
    }
    _pcsp__validate_utf8_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {5, 16},
        {247, 24},
        {251, 16},
        {252, 8},
        {253, 0},
        {527, 24},
        {531, 16},
        {532, 8},
        {534, 0},
    }
    _pcsp__value = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {495, 88},
        {499, 48},
        {500, 40},
        {502, 32},
        {504, 24},
        {506, 16},
        {508, 8},
        {509, 0},
        {1308, 88},
    }
    _pcsp__vnumber = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {803, 104},
        {807, 48},
        {808, 40},
        {810, 32},
        {812, 24},
        {814, 16},
        {816, 8},
        {817, 0},
        {1547, 104},
    }
    _pcsp__atof_eisel_lemire64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {292, 32},
        {293, 24},
        {295, 16},
        {297, 8},
        {298, 0},
        {362, 32},
    }
    _pcsp__atof_native = [][2]uint32{
        {1, 0},
        {4, 8},
        {587, 56},
        {591, 8},
        {593, 0},
    }
    _pcsp__decimal_to_f64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1144, 56},
        {1148, 48},
        {1149, 40},
        {1151, 32},
        {1153, 24},
        {1155, 16},
        {1157, 8},
        {1158, 0},
        {1169, 56},
    }
    _pcsp__right_shift = [][2]uint32{
        {1, 0},
        {318, 8},
        {319, 0},
        {387, 8},
        {388, 0},
        {396, 8},
        {398, 0},
    }
    _pcsp__left_shift = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {363, 24},
        {364, 16},
        {366, 8},
        {367, 0},
        {470, 24},
        {471, 16},
        {473, 8},
        {474, 0},
        {486, 24},
    }
    _pcsp__vsigned = [][2]uint32{
        {1, 0},
        {4, 8},
        {112, 16},
        {113, 8},
        {114, 0},
        {125, 16},
        {126, 8},
        {127, 0},
        {260, 16},
        {261, 8},
        {262, 0},
        {266, 16},
        {267, 8},
        {268, 0},
        {306, 16},
        {307, 8},
        {308, 0},
        {316, 16},
        {317, 8},
        {319, 0},
    }
    _pcsp__vstring = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {105, 56},
        {109, 40},
        {110, 32},
        {112, 24},
        {114, 16},
        {116, 8},
        {118, 0},
    }
    _pcsp__vunsigned = [][2]uint32{
        {1, 0},
        {71, 8},
        {72, 0},
        {83, 8},
        {84, 0},
        {107, 8},
        {108, 0},
        {273, 8},
        {274, 0},
        {312, 8},
        {313, 0},
        {320, 8},
        {322, 0},
    }
)

var Funcs = []loader.CFunc{
    {"__native_entry__", 0, 67, 0, nil},
    {"_f32toa", _entry__f32toa, _size__f32toa, _stack__f32toa, _pcsp__f32toa},
    {"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
    {"_format_significand", _entry__format_significand, _size__format_significand, _stack__format_significand, _pcsp__format_significand},
    {"_format_integer", _entry__format_integer, _size__format_integer, _stack__format_integer, _pcsp__format_integer},
    {"_get_by_path", _entry__get_by_path, _size__get_by_path, _stack__get_by_path, _pcsp__get_by_path},
    {"_fsm_exec", _entry__fsm_exec, _size__fsm_exec, _stack__fsm_exec, _pcsp__fsm_exec},
    {"_advance_string", _entry__advance_string, _size__advance_string, _stack__advance_string, _pcsp__advance_string},
    {"_advance_string_default", _entry__advance_string_default, _size__advance_string_default, _stack__advance_string_default, _pcsp__advance_string_default},
    {"_do_skip_number", _entry__do_skip_number, _size__do_skip_number, _stack__do_skip_number, _pcsp__do_skip_number},
    {"_skip_one_fast", _entry__skip_one_fast, _size__skip_one_fast, _stack__skip_one_fast, _pcsp__skip_one_fast},
    {"_html_escape", _entry__html_escape, _size__html_escape, _stack__html_escape, _pcsp__html_escape},
    {"_i64toa", _entry__i64toa, _size__i64toa, _stack__i64toa, _pcsp__i64toa},
    {"_u64toa", _entry__u64toa, _size__u64toa, _stack__u64toa, _pcsp__u64toa},
    {"_lspace", _entry__lspace, _size__lspace, _stack__lspace, _pcsp__lspace},
    {"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
    {"_skip_array", _entry__skip_array, _size__skip_array, _stack__skip_array, _pcsp__skip_array},
    {"_skip_number", _entry__skip_number, _size__skip_number, _stack__skip_number, _pcsp__skip_number},
    {"_skip_object", _entry__skip_object, _size__skip_object, _stack__skip_object, _pcsp__skip_object},
    {"_skip_one", _entry__skip_one, _size__skip_one, _stack__skip_one, _pcsp__skip_one},
    {"_unquote", _entry__unquote, _size__unquote, _stack__unquote, _pcsp__unquote},
    {"_validate_one", _entry__validate_one, _size__validate_one, _stack__validate_one, _pcsp__validate_one},
    {"_validate_utf8", _entry__validate_utf8, _size__validate_utf8, _stack__validate_utf8, _pcsp__validate_utf8},
    {"_validate_utf8_fast", _entry__validate_utf8_fast, _size__validate_utf8_fast, _stack__validate_utf8_fast, _pcsp__validate_utf8_fast},
    {"_value", _entry__value, _size__value, _stack__value, _pcsp__value},
    {"_vnumber", _entry__vnumber, _size__vnumber, _stack__vnumber, _pcsp__vnumber},
    {"_atof_eisel_lemire64", _entry__atof_eisel_lemire64, _size__atof_eisel_lemire64, _stack__atof_eisel_lemire64, _pcsp__atof_eisel_lemire64},
    {"_atof_native", _entry__atof_native, _size__atof_native, _stack__atof_native, _pcsp__atof_native},
    {"_decimal_to_f64", _entry__decimal_to_f64, _size__decimal_to_f64, _stack__decimal_to_f64, _pcsp__decimal_to_f64},
    {"_right_shift", _entry__right_shift, _size__right_shift, _stack__right_shift, _pcsp__right_shift},
    {"_left_shift", _entry__left_shift, _size__left_shift, _stack__left_shift, _pcsp__left_shift},
    {"_vsigned", _entry__vsigned, _size__vsigned, _stack__vsigned, _pcsp__vsigned},
    {"_vstring", _entry__vstring, _size__vstring, _stack__vstring, _pcsp__vstring},
    {"_vunsigned", _entry__vunsigned, _size__vunsigned, _stack__vunsigned, _pcsp__vunsigned},
}
