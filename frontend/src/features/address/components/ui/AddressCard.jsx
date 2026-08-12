import { Text, CustomButton } from "@/components/ui";

export default function AddressCard({ address }) {
	return (
		<div className="border-2 border-border p-2.5 md:p-5 rounded-md">
			<div className="flex justify-between mb-5">
				<Text text={address.addressLabel} className="font-medium" />
				<Text
					text={address.isDefault ? "Default" : ""}
					className="font-medium"
					textColor="text-accent-brand"
				/>
			</div>
			<div className="flex flex-col gap-2.5 mb-5">
				<Text text={address.fullName} className="font-medium" />
				<Text text={address.street} />
				<Text text={address.city} />
				<div className="flex gap-1">
					<Text text={`${address.governorate},`} />
					<Text text={address.country} />
				</div>
			</div>
			<div className="flex flex-col md:flex-row gap-2.5 justify-between">
				{!address.isDefault && (
					<CustomButton text="Set as default" onClick={() => {}} />
				)}
				<CustomButton text="Edit" onClick={() => {}} />
				<CustomButton
					text="Delete"
					onClick={() => {}}
					className="bg-destructive hover:bg-destructive-hover text-destructive-foreground hover:text-destructive-foreground-hover"
				/>
			</div>
		</div>
	);
}
