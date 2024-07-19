// exemplo de uso: <Category category={'diversity'} circleSize={'w-[32px] h-[32px]'} iconSize={20}/>

import { View } from "react-native";
import DiversityIcon from "../assets/diversityIcon.svg";
import IntegrityIcon from "../assets/integrityIcon.svg";
import HealthIcon from "../assets/healthIcon.svg";
import ConservationIcon from "../assets/conservationIcon.svg";
import ProductivityIcon from "../assets/productivityIcon.svg";
import EnvironmentalImpactIcon from "../assets/environmentaImpactIcon.svg";

export const categoryMap = {
  diversity: { icon: DiversityIcon, color: 'bg-[#C98293]' },
  integrity: { icon: IntegrityIcon, color: 'bg-[#FAA635]' },
  health: { icon: HealthIcon, color: 'bg-[#F2CF52]' },
  conservation: { icon: ConservationIcon, color: 'bg-[#3A8A88]' },
  productivity: { icon: ProductivityIcon, color: 'bg-[#A3B3FF]' },
  environmentalImpact: { icon: EnvironmentalImpactIcon, color: 'bg-[#A3BFB7]' },
};

export default function Category({ category, circleSize, iconSize }) {
  const CategoryIcon = categoryMap[category].icon;
  const color = categoryMap[category].color;

  return (
    <View className={`rounded-full ${color} ${circleSize} items-center justify-center`}>
      <CategoryIcon width={iconSize} height={iconSize} />
    </View>
  );
}
