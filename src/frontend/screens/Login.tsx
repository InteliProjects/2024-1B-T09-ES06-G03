import React from "react";
import { Text, View, TextInput, TouchableOpacity, Alert } from "react-native";
import { LinearGradient } from "expo-linear-gradient";
import { useForm, Controller } from "react-hook-form";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import Logo from "../assets/logo.svg";
import Eye from "../assets/password-eye.svg";
import { login } from "../services/api";

const schema = yup.object().shape({
  email: yup.string().email("Email inválido").required("Email é obrigatório"),
  password: yup.string().required("Senha é obrigatória"),
});

export default function Home({ navigation }) {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(schema),
  });

  const onSubmit = async (data) => {
    try {
      const token = await login(data.email, data.password);
      navigation.navigate("App");
    } catch (error) {
      Alert.alert("Erro ao fazer login", error.message);
    }
  };

  return (
    <View className="flex-1 w-full bg-gray-100">
      <View className="items-center mt-10">
        <Logo height={100} width={100} />
      </View>
      <View className="items-center mt-6">
        <Text className="text-4xl font-medium font-inter">
          Bem-vindo de volta!
        </Text>
      </View>
      <View className="items-center mt-14">
        <Text className="ml-10 self-start">E-mail</Text>
        <Controller
          control={control}
          name="email"
          render={({ field: { onChange, onBlur, value } }) => (
            <TextInput
              onBlur={onBlur}
              onChangeText={onChange}
              value={value}
              keyboardType="email-address"
              placeholder="Digite seu e-mail"
              className="border-2 border-green-10 rounded-lg p-2 w-4/5 mt-2"
            />
          )}
        />
        {/* {errors.email && <Text style={{ color: 'red' }}>{errors.email.message}</Text>} */}
      </View>
      <View className="items-center mt-12">
        <Text className="ml-10 self-start">Senha</Text>
        <View className="border-2 border-green-10 rounded-lg p-2 w-4/5 flex flex-row items-center justify-between">
          <Controller
            control={control}
            name="password"
            render={({ field: { onChange, onBlur, value } }) => (
              <TextInput
                onBlur={onBlur}
                onChangeText={onChange}
                value={value}
                placeholder="Digite sua senha"
                secureTextEntry
                className="border-none w-4/5"
              />
            )}
          />
          {/* {errors.password && <Text style={{ color: 'red' }}>{errors.password.message}</Text>} */}
          <Eye height={20} width={20} />
        </View>
        <Text className="mr-10 font-medium mt-2 text-sm self-end">
          Esqueceu a senha?
        </Text>
      </View>
      <LinearGradient
        className="w-4/5 mt-20 rounded-3xl self-center"
        colors={["rgba(20, 81, 79, 1)", "rgba(58, 138, 136, 1)"]}
        start={{ x: 0, y: 0 }}
        end={{ x: 1, y: 0 }}
      >
        <TouchableOpacity
          className="w-full justify-center items-center h-16"
          onPress={handleSubmit(onSubmit)}
        >
          <Text className="text-white font-bold text-xl font-inter">Login</Text>
        </TouchableOpacity>
      </LinearGradient>
      <View className="flex-row justify-center items-center mt-2">
        <Text className="text-center">Não tem uma conta?</Text>
        <TouchableOpacity onPress={() => navigation.navigate("Register")}>
          <Text className="font-bold ml-1 text-green-10">Criar uma conta</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}
