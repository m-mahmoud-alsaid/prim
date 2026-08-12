import SideBarTitle from "@/features/products/components/ui/sideBarTitle";
import { Checkbox } from "@/components/ui/Checkbox";

export default function SideBarDiscount() {
	const discounts = [
		{
			id: "discount-1",
			discount: "10% or more",
		},
		{
			id: "discount-2",
			discount: "25% or more",
		},
		{
			id: "discount-3",
			discount: "50% or more",
		},
	];

	return (
		<div className="border-b border-border pb-5">
			<SideBarTitle title="discount" />
			<div className="flex flex-col gap-2.5">
				{discounts.map((discount) => (
					<Checkbox key={discount.id} labelTxt={discount.discount} />
				))}
			</div>
		</div>
	);
}
