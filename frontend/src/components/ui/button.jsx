export function CustomButton({ text, onClick, className }) {
	return (
		<button
			className={`${className || "bg-primary text-primary-foreground hover:bg-accent hover:text-accent-foreground"} p-2.5 border border-border w-full h-full rounded-md font-medium text-txt-sm md:text-txt-md lg:text-txt-lg`}
			onClick={onClick}
		>
			{text}
		</button>
	);
}
