import { CustomButton } from "@/components/ui";
import { useTranslation } from "react-i18next";

export default function DangerZone() {
	const { t } = useTranslation("settings");

	const dangerActions = [
		{
			id: "danger-logout",
			buttonText: "logout",
		},
		{
			id: "danger-delete-acc",
			buttonText: "deleteAccount",
		},
	];

	return (
		<div className="flex flex-col gap-2.5 border-t border-b border-destructive border-dashed pt-5 pb-5">
			{dangerActions.map((action) => (
				<CustomButton
					key={action.id}
					text={t(`settings.${action.buttonText}`)}
					onClick={() => {}}
					className="bg-destructive text-destructive-foreground rounded-md hover:bg-destructive-hover hover:text-destructive-foreground-hover"
				/>
			))}
		</div>
	);
}
