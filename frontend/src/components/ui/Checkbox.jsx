export function Checkbox({ labelTxt, className }) {
	return (
		<label className={`flex gap-2.5 items-center group ${className}`}>
			<input type="checkbox" className="size-4" />
			<p className="text-muted-foreground group-hover:text-accent-brand">
				{labelTxt}
			</p>
		</label>
	);
}
