/*<CardSOAP
subject="Descrição detalhada do projeto..."
projectName="Way Carbon"
CEOName="Guilherme Vasconcelos"
companyName="Carbonística"
subcategory="Descarbonização"
category="productivity"
interestedNumber={10}
synergieNumber={10}
/> */

import { StatusBar } from "expo-status-bar";
import React from "react";
import { Text, View } from "react-native";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "./Card";

import { Badge } from "./Badge";

import { Avatar, AvatarImage, AvatarFallback } from "./Avatar";
import Flag from "../assets/flag.svg";
import Hands from "../assets/hands.svg";
import Category from "./Category";

export default function CardSOAP({
  subject,
  projectName,
  CEOName,
  companyName,
  subcategory,
  category,
  interestedNumber,
  synergieNumber,
}) {
  function getBadgeStyles(subCategoria) {
    switch (subCategoria) {
      case "Descarbonização":
        return {
          className: "bg-transparent border border-1 border-red-600",
          labelClasses: "text-yellow-500",
        };
      case "Sexo":
        return {
          className: "bg-transparent border border-1 border-yellow-500",
          labelClasses: "text-red-600",
        };
      default:
        return {
          className: "bg-transparent border border-1 border-gray-400",
          labelClasses: "text-black",
        };
    }
  }

  const badgeStyles = getBadgeStyles(subcategory);

  return (
    <View className="flex-1 justify-center">
      <Card>
        <CardHeader>
          <View className="flex flex-row gap-x-3">
            <Avatar>
              <AvatarImage
                source={{
                  uri: "https://pbs.twimg.com/profile_images/1745949238519803904/ZHwM5B07_400x400.jpg",
                }}
              />
              <AvatarFallback>CG</AvatarFallback>
            </Avatar>
            <CardTitle>{projectName}</CardTitle>
          </View>
          <CardDescription className="font-bold">{CEOName}</CardDescription>
          <CardDescription>{companyName}</CardDescription>
        </CardHeader>
        <CardContent className="max-h-40">
          <Text className="text-base truncate" numberOfLines={5}>
            {subject}
          </Text>
        </CardContent>
        <CardFooter className="flex justify-between">
          <View className="flex-row gap-x-2">
            <Category
              category={category}
              circleSize={"w-[32px] h-[32px]"}
              iconSize={20}
            />

            <Badge
              label={subcategory}
              className={badgeStyles.className}
              labelClasses={badgeStyles.labelClasses}
            />
          </View>
          <View className="flex-row right-0 gap-x-2">
            <View className="flex-row right-text-wrap0 gap-x-1 items-center">
              <Flag />
              <Text> {interestedNumber} </Text>
            </View>
            <View className="flex-row right-0 gap-x-1 items-center">
              <Hands />
              <Text> {synergieNumber} </Text>
            </View>
          </View>
        </CardFooter>
      </Card>
      <StatusBar style="auto" />
    </View>
  );
}
