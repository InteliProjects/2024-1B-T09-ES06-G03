import React from "react";
import { Text, View, Button, TouchableHighlight} from "react-native";
import GradientButton from "../components/GradientButton";
import LogoBlack from "../assets/LogoBlack.svg";
import Carousel from "../components/Carousel";
import ShadowButton from "../components/ShadowButton";

export default function Opening( { navigation } ) {
    return (
        <View className="w-[100%] h-[100%]">
          {/* Header */}
          <View className="self-center mt-20 h-[6%]">
            <LogoBlack/>
          </View>

          {/* Corpo */}
          <View className="flex items-center h-[65%]">
            <Carousel/>
          </View>

          {/* Footer */}
          <View className="flex w-[100%] h-[15%] items-center justify-evenly">
            <View className="w-[80%]">
              <GradientButton onPress={() => navigation.navigate('Login')} title="Login"/>
            </View>
            <View className="w-[80%]">
              <ShadowButton onPress={() => navigation.navigate('Register')} title="Cadastrar"/>
            </View>
          </View>
        </View>
      );
}