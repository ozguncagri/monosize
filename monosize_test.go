package monosize

import "testing"

const (
	OneByte      = 1
	OneKiloByte  = OneByte * 1024
	OneMegaByte  = OneKiloByte * 1024
	OneGigaByte  = OneMegaByte * 1024
	OneTeraByte  = OneGigaByte * 1024
	OnePetaByte  = OneTeraByte * 1024
	OneExaByte   = OnePetaByte * 1024
	OneZettaByte = OneExaByte * 1024
	OneYottaByte = OneZettaByte * 1024
)

func TestBytes(t *testing.T) {
	expected := "  1.00 B."
	calculated := GetFixedSize(OneByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "345.00 B."
	calculated = GetFixedSize(345 * OneByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}
}

func TestKiloBytes(t *testing.T) {
	expected := "  0.98 KB"
	calculated := GetFixedSize(OneKiloByte * 0.98)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "  1.00 KB"
	calculated = GetFixedSize(OneKiloByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "345.00 KB"
	calculated = GetFixedSize(345 * OneKiloByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}
}

func TestMegaBytes(t *testing.T) {
	expected := "  0.98 MB"
	calculated := GetFixedSize(OneMegaByte * 0.98)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "  1.00 MB"
	calculated = GetFixedSize(OneMegaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "345.00 MB"
	calculated = GetFixedSize(345 * OneMegaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}
}

func TestGigaBytes(t *testing.T) {
	expected := "  0.98 GB"
	calculated := GetFixedSize(OneGigaByte * 0.98)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "  1.00 GB"
	calculated = GetFixedSize(OneGigaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "345.00 GB"
	calculated = GetFixedSize(345 * OneGigaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}
}

func TestTeraBytes(t *testing.T) {
	expected := "  0.98 TB"
	calculated := GetFixedSize(OneTeraByte * 0.98)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "  1.00 TB"
	calculated = GetFixedSize(OneTeraByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "345.00 TB"
	calculated = GetFixedSize(345 * OneTeraByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}
}

func TestPetaBytes(t *testing.T) {
	expected := "  0.98 PB"
	calculated := GetFixedSize(OnePetaByte * 0.98)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "  1.00 PB"
	calculated = GetFixedSize(OnePetaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "345.00 PB"
	calculated = GetFixedSize(345 * OnePetaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}
}

func TestExaBytes(t *testing.T) {
	expected := "  0.98 EB"
	calculated := GetFixedSize(OneExaByte * 0.98)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "  1.00 EB"
	calculated = GetFixedSize(OneExaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "345.00 EB"
	calculated = GetFixedSize(345 * OneExaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}
}

func TestZettaBytes(t *testing.T) {
	expected := "  0.98 ZB"
	calculated := GetFixedSize(OneZettaByte * 0.98)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "  1.00 ZB"
	calculated = GetFixedSize(OneZettaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "345.00 ZB"
	calculated = GetFixedSize(345 * OneZettaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}
}

func TestYottaBytes(t *testing.T) {
	expected := "  0.98 YB"
	calculated := GetFixedSize(OneYottaByte * 0.98)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "  1.00 YB"
	calculated = GetFixedSize(OneYottaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}

	expected = "345.00 YB"
	calculated = GetFixedSize(345 * OneYottaByte)
	if expected != calculated {
		t.Errorf(`Expected "%v", got "%v"`, expected, calculated)
	}
}

func TestOutputLength(t *testing.T) {
	calculated := GetFixedSize(345 * OneYottaByte)
	if len(calculated) != 9 {
		t.Errorf(`Expected output limit is exceeded! Output length is %v`, calculated)
	}
}

func TestNegativeOrZeroInput(t *testing.T) {
	calculated := GetFixedSize(-1)
	if calculated != "  0.00 B." {
		t.Errorf(`Expected "  0.00 B.", output is %v`, calculated)
	}

	calculated = GetFixedSize(0)
	if calculated != "  0.00 B." {
		t.Errorf(`Expected "  0.00 B.", output is %v`, calculated)
	}
}
