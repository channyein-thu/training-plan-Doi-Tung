function PlanDetails({ title, subtitle }: { title: string; subtitle: string }) {
  return (
    <div>
      {" "}
      <div>
        <p className="text-sm text-gray-600">{title}</p>
        <p className="text-md text-black">{subtitle}</p>
      </div>
    </div>
  );
}

export default PlanDetails;
