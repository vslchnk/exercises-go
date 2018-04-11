package weightconv

func KToP(k KG) Pound { return Pound(k / 0.45) }

func PToK(p Pound) KG { return KG(p / 2.2) }
