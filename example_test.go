package golangvectorspace_test

import (
	"fmt"

	"github.com/go-dedup/golangvectorspace"
)

// to show the full code in GoDoc
type dummy struct {
}

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example_output() {
	var docs = []string{
		(" 2012 Ford Mustang Premium Sport Coupe Selling my rare 2012 Club of America Edition Mustang. The car is fully loaded with all options available. Car is very well maintained, always stored in garage year round and comes with a separate set… 75,000km | Automatic"),
		(" 2012 Ford Mustang Premium Sport Coupe (2 door) Selling my rare 2012 Club of America Edition Mustang. The car is fully loaded with all options available. Car is very well maintained, always stored in garage year round and comes with a separate set… 75,000km | Automatic"),
	}
	// Code starts
	for ii := range docs {
		for jj := range docs {
			concordance1 := golangvectorspace.BuildConcordance(docs[ii])
			concordance2 := golangvectorspace.BuildConcordance(docs[jj])
			r := golangvectorspace.Relation(concordance1, concordance2)
			fmt.Printf("[%d, %d] '%f'\n", ii, jj, r)
		}
	}
	// Code ends

	docs = []string{
		(" Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic"),
		(" 2016 Ford F-150 Pickup Truck Jason canopy 2016 5.1/2 Comes with 2 keys Inside carpet and light No damage Excellent condition Call 10,000km | Automatic"),
		(" Like New 2015 Ford F-150 XTR Pickup Truck 2015 Ford F-150 XTR Supercab, white with tan interior. Truck was lady owned and driven and meticulously maintained with only 27000 kms. The truck is the closest you will find to a brand new vehicle.… 27,000km | Automatic"),
		(" 2010 Ford F-150 SuperCrew Pickup Truck, 5.4L 4x4 2010 F-150 4x4 Supercrew XLT, 5-1/2ft box with spray in bed liner and soft tonneau cover, 5.4L V8 engine, XTR pkg, tow pkg with built in brake controller, Very clean. located in Orangeville. Call… 81,295km | Automatic"),
		(" Trade 2016 Ford F-150 chrome bumpers for your black ones I'm looking to trade my chrome bumpers and grill of your black bumpers and grill. Please call or text 1,234km | Automatic"),
		(" 2013 Ford F-150 XLT/XTR 4x4 3.5 Ecobust 2013 Ford F-150 XLT/XTR 4x4 3.5 Ecobust with 90000 km on it I'm a first owner. Always serviced at Dixie Ford Mississauga. Still on the warranty for 100k km or 5 years. Factory Max Tow package.… 90,000km | Automatic"),
		(" Ford F-150 EcoBoost 2014 crew cab with box liner and cover. Low kilometers. Back up camera and sensors. Satellite radio and CD player cold AC and rubber floor mats. Cloth seats. Six seater, remaining bumper to bumper… 26,000km | Automatic"),
		(" 2014 Ford F-150 XL Pickup Truck sold by owner 2014 Ford F-150 XL Pickup Truck sold by owner, low mileage (64000km). New tires, cap, hitch, no accidents, great conditions. First come first buy. Please, contact me 64,000km | Automatic | CarProof"),
		(" 2015 Ford F-150 Lariat SuperCrew Pickup Truck 2015 Ford F150 Lariat. Loaded. No sunroof. 500A Available August 1st, 2017 Presently in northern Ontario. Contact Martin 55,000km | Automatic"),
		(" 2015 Ford F-150 Lariat SuperCrew Pickup Truck 2015 Ford F150 Lariat. Loaded. No sunroof. 500A Available August 1st, 2017 Presently in northern Ontario. Contact Martin 55,000km | Automatic"),
		(" 2015 Ford F-150 SuperCrew FX4 Off Road sport Pickup Truck Mint condition 3.5 Eco Boost V6. Aluminum body. Still under warranty. Heated power Lombard seats. Brake pedal adjustment. Power windows, locks, mirrors, sunroof. Tilt and telescopic steering wheel.… 21,000km | Automatic"),
		(" 2016 Ford F-150 Lariat Pickup Truck 2.7L 501A w/FX4 Offroad Pky Vehicle Features: 2016 Magnetic Grey Ford F-150 SuperCrew 145 Wheelbase FX4 Off Road Package w/Skid Plates, 3.73 Ratio E-lock 2.7L EcoBoost Twin-Turbo Engine 6-Speed SelectShift Transmission 6500#… 19,999km | Automatic"),
		(" 2013 Ford F-150 SuperCrew Pickup Truck Great Clean truck, 301A package, Brakes just done, power mirrors, Pedals, Windows, Seats Bluetooth and satellite radio Trailer brake Controller and Class III hitch 80000 km 6.5 foot box 5.0L V8 motor… 80,000km | Automatic"),
		(" -- URGENT-- 2010 Ford F-150 XL -- MINT -- BULLET PROOF V8 — Hello, I have my beautiful Grey Ford F150 XL for sale. This truck has been driven by a senior to the grocery store and back. This truck has never been smoked in and is very clean base model. It has a… 41,500km | Automatic"),
	}

	// Code starts
	fmt.Println("================")
	for ii := range docs {
		for jj := ii + 1; jj < len(docs); jj++ {
			concordance1 := golangvectorspace.BuildConcordance(docs[ii])
			concordance2 := golangvectorspace.BuildConcordance(docs[jj])
			r := golangvectorspace.Relation(concordance1, concordance2)
			if r > 0.35 {
				fmt.Printf("[%d, %d] '%f'\n", ii, jj, r)
			}
		}
	}

	// The vector space similarity calculation algorithm appears to be
	// not symmetric but it actually is. I.e.,
	// Relation(a,b) == Relation(b,a)
	fmt.Println("================")
	for ii := range docs {
		concordance0 := golangvectorspace.BuildConcordance(docs[ii])
		// break point
		bp := wordBreak(docs[ii], len(docs[ii])/2)
		concordance1 := golangvectorspace.BuildConcordance(docs[ii][:bp])
		concordance2 := golangvectorspace.BuildConcordance(docs[ii][bp:])
		fmt.Printf("%f, %f\n%f, %f\n\n",
			golangvectorspace.Relation(concordance0, concordance1),
			golangvectorspace.Relation(concordance0, concordance2),
			golangvectorspace.Relation(concordance1, concordance0),
			golangvectorspace.Relation(concordance2, concordance0),
		)
	}
	// Code ends

	// Output:
	// [0, 0] '1.000000'
	// [0, 1] '0.981650'
	// [1, 0] '0.981650'
	// [1, 1] '1.000000'
	// ================
	// [0, 2] '0.367687'
	// [0, 13] '0.458706'
	// [2, 8] '0.356260'
	// [2, 9] '0.356260'
	// [7, 8] '0.353726'
	// [7, 9] '0.353726'
	// [8, 9] '1.000000'
	// ================
	// 0.789922, 0.752369
	// 0.789922, 0.752369
	//
	// 0.755929, 0.681385
	// 0.755929, 0.681385
	//
	// 0.805194, 0.731898
	// 0.805194, 0.731898
	//
	// 0.795495, 0.725000
	// 0.795495, 0.725000
	//
	// 0.779350, 0.779350
	// 0.779350, 0.779350
	//
	// 0.828918, 0.677759
	// 0.828918, 0.677759
	//
	// 0.773325, 0.773325
	// 0.773325, 0.773325
	//
	// 0.793116, 0.622171
	// 0.793116, 0.622171
	//
	// 0.778499, 0.651339
	// 0.778499, 0.651339
	//
	// 0.778499, 0.651339
	// 0.778499, 0.651339
	//
	// 0.722999, 0.707107
	// 0.722999, 0.707107
	//
	// 0.787839, 0.664078
	// 0.787839, 0.664078
	//
	// 0.674200, 0.753778
	// 0.674200, 0.753778
	//
	// 0.778061, 0.738615
	// 0.778061, 0.738615
}

func wordBreak(str string, breakAt int) int {
	for {
		if str[breakAt] == ' ' {
			break
		}
		breakAt--
	}
	return breakAt
}
