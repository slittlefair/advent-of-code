package main

import (
	"testing"
)

func Test_validateYr(t *testing.T) {
	type args struct {
		val string
		min int
		max int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "returns an error if the passed val can't be converted into an int",
			args: args{
				val: "",
				min: 0,
				max: 1,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "returns false if the passed val is less than the min value",
			args: args{
				val: "3",
				min: 5,
				max: 10,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "returns false if the passed val is greater than the max value",
			args: args{
				val: "13",
				min: 5,
				max: 10,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "returns true if the passed val is equal to the min value",
			args: args{
				val: "5",
				min: 5,
				max: 10,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "returns true if the passed val is equal to the max value",
			args: args{
				val: "10",
				min: 5,
				max: 10,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "returns true if the passed val is greater than the min value and less than the max value",
			args: args{
				val: "7",
				min: 5,
				max: 10,
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateYr(tt.args.val, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateYr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateYr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateHgt(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{
			name: "returns false if no matches can be found in the given val",
			val:  "some random string",
			want: false,
		},
		{
			name: "returns false if given val contains an int not followed by cm or in",
			val:  "18900mm",
			want: false,
		},
		{
			name: "returns false if given val contains cm or in and is not preceeded by an int",
			val:  "159~cm",
			want: false,
		},
		{
			name: "returns false if cm is less than 150",
			val:  "149cm",
			want: false,
		},
		{
			name: "returns false if cm is greater than 193",
			val:  "194cm",
			want: false,
		},
		{
			name: "returns true if cm is equal to 150 min",
			val:  "150cm",
			want: true,
		},
		{
			name: "returns true if cm is equal to 193 max",
			val:  "193cm",
			want: true,
		},
		{
			name: "returns true if cm is greater than 150 and less than 193",
			val:  "170cm",
			want: true,
		},
		{
			name: "returns false if in is less than 59",
			val:  "58in",
			want: false,
		},
		{
			name: "returns false if in is greater than 76",
			val:  "77in",
			want: false,
		},
		{
			name: "returns true if in is equal to 59 min",
			val:  "59in",
			want: true,
		},
		{
			name: "returns true if in is equal to 76 max",
			val:  "76in",
			want: true,
		},
		{
			name: "returns true if cm is greater than 59 and less than 76",
			val:  "67in",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHgt(tt.val); got != tt.want {
				t.Errorf("validateHgt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateHcl(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{
			name: "returns false if the given value doesn't begin with #",
			val:  "f65193",
			want: false,
		},
		{
			name: "returns false if the given value's hex string is less than 6 characters",
			val:  "#f653a",
			want: false,
		},
		{
			name: "returns false if the given value's hex string is greater than 6 characters",
			val:  "#f65193a",
			want: false,
		},
		{
			name: "returns false if the given value contains a non alphanumeric character",
			val:  "#4562.9",
			want: false,
		},
		{
			name: "returns false if the given value contains a character outside of hex range",
			val:  "#45s629",
			want: false,
		},
		{
			name: "returns true if the given value is a valid hex string",
			val:  "#456290",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHcl(tt.val); got != tt.want {
				t.Errorf("validateHcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateEcl(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{
			name: "returns true if the given val is amb",
			val:  "amb",
			want: true,
		},
		{
			name: "returns true if the given val is blu",
			val:  "blu",
			want: true,
		},
		{
			name: "returns true if the given val is brn",
			val:  "brn",
			want: true,
		},
		{
			name: "returns true if the given val is gry",
			val:  "gry",
			want: true,
		},
		{
			name: "returns true if the given val is grn",
			val:  "grn",
			want: true,
		},
		{
			name: "returns true if the given val is hzl",
			val:  "hzl",
			want: true,
		},
		{
			name: "returns true if the given val is oth",
			val:  "oth",
			want: true,
		},
		{
			name: "returns false if the given val is not one of those described",
			val:  "sloths",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateEcl(tt.val); got != tt.want {
				t.Errorf("validateEcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatePid(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{
			name: "returns false if the val contains an invalid character",
			val:  "0984t6712",
			want: false,
		},
		{
			name: "returns false if the val is less than 9 characters",
			val:  "09846712",
			want: false,
		},
		{
			name: "returns false if the val is greater than 9 characters",
			val:  "0984671892",
			want: false,
		},
		{
			name: "returns true if the val is 9 ints",
			val:  "000009145",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePid(tt.val); got != tt.want {
				t.Errorf("validatePid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allFieldsValid(t *testing.T) {
	tests := []struct {
		name    string
		fields  map[string]string
		want    bool
		wantErr bool
	}{
		{
			name: "returns false if byr field does not pass validation",
			fields: map[string]string{
				"byr": "199",
				"iyr": "2015",
				"eyr": "2025",
				"hgt": "180cm",
				"hcl": "#e89af4",
				"ecl": "blu",
				"pid": "098765432",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "returns an error if byr field is not an int",
			fields: map[string]string{
				"byr": "?",
				"iyr": "2015",
				"eyr": "2025",
				"hgt": "180cm",
				"hcl": "#e89af4",
				"ecl": "blu",
				"pid": "098765432",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "returns false if iyr field does not pass validation",
			fields: map[string]string{
				"byr": "1991",
				"iyr": "201",
				"eyr": "2025",
				"hgt": "180cm",
				"hcl": "#e89af4",
				"ecl": "blu",
				"pid": "098765432",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "returns an error if iyr field is not an int",
			fields: map[string]string{
				"byr": "1991",
				"iyr": "*",
				"eyr": "2025",
				"hgt": "180cm",
				"hcl": "#e89af4",
				"ecl": "blu",
				"pid": "098765432",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "returns false if eyr field does not pass validation",
			fields: map[string]string{
				"byr": "1991",
				"iyr": "2015",
				"eyr": "20258",
				"hgt": "180cm",
				"hcl": "#e89af4",
				"ecl": "blu",
				"pid": "098765432",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "returns an error if eyr field is not an int",
			fields: map[string]string{
				"byr": "1991",
				"iyr": "2015",
				"eyr": "(!)",
				"hgt": "180cm",
				"hcl": "#e89af4",
				"ecl": "blu",
				"pid": "098765432",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "returns false if hgt field does not pass validation",
			fields: map[string]string{
				"byr": "1991",
				"iyr": "2015",
				"eyr": "2025",
				"hgt": "18cm",
				"hcl": "#e89af4",
				"ecl": "blu",
				"pid": "098765432",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "returns false if hcl field does not pass validation",
			fields: map[string]string{
				"byr": "1991",
				"iyr": "2015",
				"eyr": "2025",
				"hgt": "180cm",
				"hcl": "e89af4",
				"ecl": "blu",
				"pid": "098765432",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "returns false if ecl field does not pass validation",
			fields: map[string]string{
				"byr": "1991",
				"iyr": "2015",
				"eyr": "2025",
				"hgt": "180cm",
				"hcl": "#e89af4",
				"ecl": "blue",
				"pid": "098765432",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "returns false if pid field does not pass validation",
			fields: map[string]string{
				"byr": "1991",
				"iyr": "2015",
				"eyr": "2025",
				"hgt": "180cm",
				"hcl": "#e89af4",
				"ecl": "blu",
				"pid": "0987654321",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "returns true if all fields are valid",
			fields: map[string]string{
				"byr": "1991",
				"iyr": "2015",
				"eyr": "2025",
				"hgt": "180cm",
				"hcl": "#e89af4",
				"ecl": "blu",
				"pid": "098765432",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := allFieldsValid(tt.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("allFieldsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("allFieldsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution(t *testing.T) {
	tests := []struct {
		name    string
		entries []string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "advent of code part 1 example",
			entries: []string{
				"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
				"byr:1937 iyr:2017 cid:147 hgt:183cm",
				"",
				"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
				"hcl:#cfa07d byr:1929",
				"",
				"hcl:#ae17e1 iyr:2013",
				"eyr:2024",
				"ecl:brn pid:760753108 byr:1931",
				"hgt:179cm",
				"",
				"hcl:#cfa07d eyr:2025 pid:166559648",
				"iyr:2011 ecl:brn hgt:59in",
			},
			want:    2,
			want1:   2,
			wantErr: false,
		},
		{
			name: "advent of code part 2 example all invalid",
			entries: []string{
				"eyr:1972 cid:100",
				"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
				"",
				"iyr:2019",
				"hcl:#602927 eyr:1967 hgt:170cm",
				"ecl:grn pid:012533040 byr:1946",
				"",
				"hcl:dab227 iyr:2012",
				"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
				"",
				"hgt:59cm ecl:zzz",
				"eyr:2038 hcl:74454a iyr:2023",
				"pid:3556412378 byr:2007",
			},
			want:    4,
			want1:   0,
			wantErr: false,
		},
		{
			name: "advent of code part 2 example all valid",
			entries: []string{
				"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
				"hcl:#623a2f",
				"",
				"eyr:2029 ecl:blu cid:129 byr:1989",
				"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
				"",
				"hcl:#888785",
				"hgt:164cm byr:2001 iyr:2015 cid:88",
				"pid:545766238 ecl:hzl",
				"eyr:2022",
				"",
				"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			},
			want:    4,
			want1:   4,
			wantErr: false,
		},
		{
			name: "returns an error if one of the fields returns an error",
			entries: []string{
				"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
				"hcl:#623a2f",
				"",
				"eyr:2029 ecl:blu cid:129 byr:1989",
				"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
				"",
				"hcl:#888785",
				"hgt:164cm byr:2001 iyr:20?15 cid:88",
				"pid:545766238 ecl:hzl",
				"eyr:2022",
				"",
				"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := solution(tt.entries)
			if (err != nil) != tt.wantErr {
				t.Errorf("solution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solution() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("solution() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
