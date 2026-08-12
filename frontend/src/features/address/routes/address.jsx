import { useState } from "react";
import AddressLayout from "@/features/address/components/layout/addressLayout";

export function Address() {
	const [addresses, setAddresses] = useState([
		{
			id: "address-one",
			addressLabel: "Home",
			isDefault: true,
			fullName: "Mohamed Mahmoud",
			street: "123 example street",
			city: "Mit Ghamr",
			governorate: "Dakahlya",
			country: "Egypt",
		},
		{
			id: "address-two",
			addressLabel: "Work",
			isDefault: false,
			fullName: "Mohamed Mahmoud",
			street: "123 example street",
			city: "Mit Ghamr",
			governorate: "Dakahlya",
			country: "Egypt",
		},
		{
			id: "address-three",
			addressLabel: "Other",
			isDefault: false,
			fullName: "Mohamed Mahmoud",
			street: "123 example street",
			city: "Mit Ghamr",
			governorate: "Dakahlya",
			country: "Egypt",
		},
	]);

	const addAddress = (
		id,
		addressLabel,
		isDefault,
		fullName,
		street,
		city,
		governorate,
		country,
	) =>
		setAddresses((prev) => [
			...prev,
			{
				id: id,
				addressLabel: addressLabel,
				isDefault: isDefault,
				fullName: fullName,
				street: street,
				city: city,
				governorate: governorate,
				country: country,
			},
		]);

	return <AddressLayout addresses={addresses} addAddress={addAddress} />;
}
