import React from "react";

function PlanDetails({ params }: { params: { planId: string } }) {
  return <div>PlanDetails for plan ID: {params.planId}</div>;
}

export default PlanDetails;
