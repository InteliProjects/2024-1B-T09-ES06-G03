import React from "react";
import { Text, View, TextInput, TouchableOpacity } from "react-native";
import { LinearGradient } from 'expo-linear-gradient';
import { useForm, Controller } from 'react-hook-form';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import Logo from '../assets/logo.svg';
import  InputComponent  from '../components/Input';

const schema = yup.object().shape({
    email: yup.string().email('Email inválido').required('Email é obrigatório'),
    confirmEmail: yup.string().oneOf([yup.ref('email'), null], 'Os emails devem ser iguais').required('Confirmação de email é obrigatória')
});

export default function RegisterEmail({ route, navigation }) {
    const { control, handleSubmit, formState: { errors } } = useForm({
        defaultValues: route.params.formData,
        resolver: yupResolver(schema)
      });

const onNext = (data) => {
    navigation.navigate('CadastroSenha', { formData: data });
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
          <Text className="text-2xl font-medium font-inter">Informações Pessoais</Text>
        </View>
        <View className="items-center mt-10">
        <Text className="ml-10 self-start">Email</Text>
        <Controller
        control={control}
        name="email"
        render={({ field: { onChange, onBlur, value } }) => (
          <TextInput
          onBlur={onBlur}
          onChangeText={onChange}
          keyboardType="email-address"
          value={value}
          placeholder="Digite seu email"
          className="border-2 border-green-10 rounded-lg p-2 w-4/5 mt-2"
        />
        )}
      />
        </View>
        <View className="items-center mt-10">
        <Text className="ml-10 self-start">Confirme seu email</Text>
        <Controller
        control={control}
        name="confirmEmail"
        render={({ field: { onChange, onBlur, value } }) => (
          <TextInput
          onBlur={onBlur}
          onChangeText={onChange}
          keyboardType="email-address"
          value={value}
          placeholder="Confirme seu email"
          className="border-2 border-green-10 rounded-lg p-2 w-4/5 mt-2"
        />
        )}
      />
        </View>
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
          <Text className="text-white font-bold text-xl font-inter">Continuar</Text>
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