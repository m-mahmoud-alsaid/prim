import { Title, Text } from "@/components/ui";
import AddAddress from "@/features/address/components/ui/AddAddress";
import AddressesGrid from "@/features/address/components/ui/AddressesGrid";

export default function AddressLayout({ addresses, addAddress }) {
	return (
		<div className="">
			<Title
				title="Addresses"
				subtitle="Manage your saved shipping addresses."
			/>
			{addresses.length === 0 && (
				<Text
					text={`You haven't added any addresses yet. 
                    Add an address to make checkout faster.`}
					className="text-muted-foreground mb-5"
				/>
			)}
			<AddAddress addAddress={addAddress} />

			{addresses.length > 0 && <AddressesGrid addresses={addresses} />}
		</div>
	);
}
