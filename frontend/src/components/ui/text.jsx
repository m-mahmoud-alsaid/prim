export function Text({ text, className, textColor }) {
	return (
		<p
			className={`${className} ${textColor || "text-foreground"} text-txt-sm md:text-txt-md lg:text-txt-lg`}
		>
			{text}
		</p>
	);
}
