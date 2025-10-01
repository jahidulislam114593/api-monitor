import React from "react";

const StatCard = ({
  icon: Icon,
  title,
  value,
  bgColor,
  textColor,
  iconColor,
}) => {
  return (
    <div className={`${bgColor} p-4 rounded-lg`}>
      <div className="flex items-center gap-2">
        <Icon size={20} className={iconColor} />
        <div>
          <div className={`text-2xl font-bold ${textColor}`}>
            {value || "0"}
          </div>
          <div className={`text-sm ${textColor.replace("600", "700")}`}>
            {title}
          </div>
        </div>
      </div>
    </div>
  );
};

export default StatCard;
