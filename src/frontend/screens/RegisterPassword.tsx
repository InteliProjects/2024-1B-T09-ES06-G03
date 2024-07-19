import React from "react";
import { Text, View, TextInput, TouchableOpacity } from "react-native";
import { LinearGradient } from 'expo-linear-gradient';
import { useForm, Controller } from 'react-hook-form';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import Logo from '../assets/logo.svg';
import  InputComponent  from '../components/Input';
import Eye from '../assets/password-eye.svg';

const schema = yup.object().shape({
    password: yup.string().required('Senha é obrigatória'),
    confirmPassword: yup.string().oneOf([yup.ref('password'), null], 'As senhas devem ser iguais').required('Confirmação de senha é obrigatória')
});

export default function RegisterPassword({ route, navigation }) {
    const { control, handleSubmit, formState: { errors } } = useForm({
        defaultValues: route.params.formData,
        resolver: yupResolver(schema)
      });


const onNext = (data) => {
    console.log(data);
    navigation.navigate('App')
}
    return (
     <View className="flex-1 w-full bg-gray-100">
        <View className="items-center mt-5">
          <Logo height={100} width={100} />
        </View>
        <View className="items-center mt-4">
          <Text className="text-4xl font-medium font-inter">Seja bem-vindo!</Text>
        </View>
        <View className="items-center mt-8">
          <Text className="text-2xl font-medium font-inter">Inserir senha</Text>
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
                placeholder="Crie uma senha"
                secureTextEntry
                className="border-none w-4/5"
              />
            )}
          />
          {/* {errors.password && <Text style={{ color: 'red' }}>{errors.password.message}</Text>} */}
          <Eye height={20} width={20} />
        </View>
    </ View>
    <View className="items-center mt-12">
        <Text className="ml-10 self-start">Confirme sua senha</Text>
        <View className="border-2 border-green-10 rounded-lg p-2 w-4/5 flex flex-row items-center justify-between">
          <Controller
            control={control}
            name="confirmPassword"
            render={({ field: { onChange, onBlur, value } }) => (
              <TextInput
                onBlur={onBlur}
                onChangeText={onChange}
                value={value}
                placeholder="Confirme sua senha"
                secureTextEntry
                className="border-none w-4/5"
              />
            )}
          />
          {/* {errors.password && <Text style={{ color: 'red' }}>{errors.password.message}</Text>} */}
          <Eye height={20} width={20} />
        </View>
    </ View>
    <LinearGradient
        className="w-4/5 mt-20 rounded-3xl self-center"
        colors={['rgba(20, 81, 79, 1)', 'rgba(58, 138, 136, 1)']}
        start={{ x: 0, y: 0 }}
        end={{ x: 1, y: 0 }}
      >
        <TouchableOpacity
          className="w-full justify-center items-center h-16"
          onPress={handleSubmit(onNext)}
        >
          <Text className="text-white font-bold text-xl font-inter">Finalizar cadastro</Text>
        </TouchableOpacity>
      </LinearGradient>
      <View className="flex-row justify-center items-center mt-2">
        <Text className="text-center">Já tem uma conta?</Text>
        <TouchableOpacity onPress={() => navigation.navigate('Login')}>
          <Text className="font-bold ml-1 text-green-10">Fazer login</Text>
        </TouchableOpacity>
      </View>
    </ View>
    )
}