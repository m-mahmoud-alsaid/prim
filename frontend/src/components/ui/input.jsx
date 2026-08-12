export function CustomInput({
	type,
	value,
	isDisabled,
	placeholder,
	onChange,
	className,
}) {
	return (
		<input
			type={type}
			placeholder={placeholder}
			onChange={onChange}
			value={value}
			disabled={isDisabled}
			className={`${className} p-2.5 truncate text-txt-sm md:text-txt-md lg:text-txt-lg placeholder:text-muted-foreground disabled:text-muted-foreground`}
		/>
	);
}
