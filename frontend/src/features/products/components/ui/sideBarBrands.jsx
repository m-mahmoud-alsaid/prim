import SideBarTitle from "@/features/products/components/ui/sideBarTitle";
import { Checkbox } from "@/components/ui/Checkbox";

export default function SideBarBrands() {
	const brands = [
		{
			id: "brand-1",
			brand: "Sony",
		},
		{
			id: "brand-2",
			brand: "JBL",
		},
		{
			id: "brand-3",
			brand: "Sumsang",
		},
		{
			id: "brand-4",
			brand: "Apple",
		},
	];

	return (
		<div className="border-b border-border pb-5">
			<SideBarTitle title="brands" />
			<div className="flex flex-col gap-2.5">
				{brands.map((brand) => (
					<Checkbox key={brand.id} labelTxt={brand.brand} />
				))}
			</div>
			<p className="mt-2.5 text-accent-brand hover:underline hover:underline-offset-2 cursor-pointer font-medium">
				Show more
			</p>
		</div>
	);
}
