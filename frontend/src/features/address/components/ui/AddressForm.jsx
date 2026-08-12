import { CustomButton, CustomInput, Checkbox, Text } from "@/components/ui";

export default function AddressForm({ add, addAddress }) {
	const formFields = [
		{
			id: "fullName-12ksdf",
			labelTitle: "Full Name",
			input: {
				type: "text",
				placeholder: "Enter your full name",
				isRequired: true,
				onChange: () => {},
			},
		},
		{
			id: "phone-12ksdf",
			labelTitle: "Phone",
			input: {
				type: "number",
				placeholder: "Enter your phone",
				isRequired: true,
				onChange: () => {},
			},
		},
		{
			id: "street-address-12ksdf",
			labelTitle: "Street Address",
			input: {
				type: "text",
				placeholder: "Enter your street address",
				isRequired: true,
				onChange: () => {},
			},
		},
		{
			id: "apartment-12ksdf",
			labelTitle: "Apartment / Unit",
			input: {
				type: "text",
				placeholder: "Enter your apartment address",
				isRequired: false,
				onChange: () => {},
			},
		},
		{
			id: "country-12ksdf",
			labelTitle: "Country",
			input: {
				type: "text",
				placeholder: "Enter your country",
				isRequired: true,
				onChange: () => {},
			},
		},
		{
			id: "governorate-12ksdf",
			labelTitle: "Governorate",
			input: {
				type: "text",
				placeholder: "Enter your governorate",
				isRequired: true,
				onChange: () => {},
			},
		},
		{
			id: "city-12ksdf",
			labelTitle: "City",
			input: {
				type: "text",
				placeholder: "Enter your city",
				isRequired: true,
				onChange: () => {},
			},
		},
	];

	return (
		<form className="" onSubmit={(e) => e.preventDefault()}>
			<div className="flex flex-col gap-2.5 mb-5">
				{formFields.map((field) => (
					<label className="" key={field.id}>
						<Text
							text={field.labelTitle}
							className={`mb-1 relative w-fit${field.input.isRequired ? "after:absolute after:content-['*'] after:text-destructive after:ml-1 after:top-0" : ""}`}
						/>
						<CustomInput
							type={field.input.type}
							placeholder={field.input.placeholder}
							onChange={field.input.onChange}
							className="w-full bg-input-background rounded-sm"
						/>
					</label>
				))}
			</div>
			<Checkbox labelTxt="Set as default address" className="mb-5" />
			<div className="flex flex-col md:flex-row gap-2.5">
				<CustomButton
					text="Cancel"
					onClick={add}
					className="bg-destructive hover:bg-destructive-hover text-destructive-foreground hover:text-destructive-foreground-hover"
				/>
				<CustomButton
					text="Save Address"
					onClick={addAddress}
					className=""
				/>
			</div>
		</form>
	);
}
