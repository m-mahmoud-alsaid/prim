import AddressCard from "@/features/address/components/ui/AddressCard";

export default function AddressesGrid({ addresses }) {
	return (
		<div className="grid grid-cols-1 xl:grid-cols-2 gap-2.5 mt-5">
			{addresses.map((address) => (
				<AddressCard key={address.id} address={address} />
			))}
		</div>
	);
}
