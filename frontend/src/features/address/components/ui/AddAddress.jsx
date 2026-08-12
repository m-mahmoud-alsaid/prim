import { useState } from "react";
import { Text, CustomButton } from "@/components/ui";
import AddressForm from "@/features/address/components/ui/AddressForm";

export default function AddAddress({ addAddress }) {
	const [isAdd, setIsAdd] = useState(false);
	const add = () => setIsAdd((prev) => !prev);

	return (
		<div className="">
			<CustomButton text="Add New Address" onClick={add} />
			{isAdd && (
				<div className="bg-popover text-popover-foreground p-5 rounded-md border-2 border-border">
					<Text text="Add New Address" className="font-medium mb-5" />
					<AddressForm add={add} addAddress={addAddress} />
				</div>
			)}
		</div>
	);
}
