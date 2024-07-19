import React from "react";
import { Text, View, StyleSheet, TouchableOpacity, Image, Dimensions } from "react-native";
import InputComponent from '../components/Input';
import { KeyboardAwareScrollView } from 'react-native-keyboard-aware-scroll-view';
import GradientButton from '../components/GradientButton';
import SocialInput from '../components/SocialInput';

const { width } = Dimensions.get('window');

const usersData = {
  name: "John Doe",
  company: "Tech Solutions",
  description: "This is a sample description for the project. It includes detailed information about the project's scope, goals, and other relevant details.",
  social: {
    instagram: "johndoe_insta",
    linkedin: "johndoe_linkedin"
  }
};

export default function EditProfile({ navigation }) {
  return (
    <KeyboardAwareScrollView
      style={{ flex: 1, backgroundColor: '#fff'}}
      scrollEnabled={true}
    >
      <View className="justify-center items-center">
        <View style={styles.container}>
          <TouchableOpacity style={styles.imageUpload}>
            <Image source={require('../assets/imageIcon.png')} style={styles.uploadIcon} />
            <Image source={require('../assets/uploadIcon.png')} style={styles.smallUploadIcon} />
          </TouchableOpacity>
          <InputComponent
            label="Nome"
            type="input"
            placeholder={usersData.name}
            maxLength={50}
            onChangeText={''}
            value={''}
          />
          <InputComponent
            label="Empresa"
            type="input"
            placeholder={usersData.company}
            maxLength={50}
            onChangeText={''}
            value={''}
          />
        </View>
        <View style={styles.socialContainer}>
          <Text style={styles.socialLabel}>
            Redes Sociais
          </Text>
          <View>
            <SocialInput type="Instagram" placeholder={usersData.social.instagram} />
            <SocialInput type="LinkedIn" placeholder={usersData.social.linkedin} />
          </View>
        </View>
        <View style={styles.descriptionContainer}>
          <InputComponent
            label="Descrição"
            type="textarea"
            placeholder={usersData.description}
            maxLength={500}
            onChangeText={''}
            value={''}
          />
        </View>
        <View className="w-[90%] mt-4 items-center justify-center bg-purple-200">
          <GradientButton title="Salvar Alterações" onPress={() => navigation.navigate('Perfil')} />
        </View>
      </View>
    </KeyboardAwareScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    justifyContent: 'center',
    backgroundColor: '#fff',
    alignItems: 'center',
  },
  input: {
    width: width * 0.9,
    marginBottom: width * 0.01,
    backgroundColor: '#fff',
    borderRadius: 5,
    zIndex: 1,
  },
  dropdown: {
    width: width * 0.9,
    marginBottom: width * 0.01,
    backgroundColor: '#fff',
    borderRadius: 5,
    zIndex: 1,
  },
  button: {
    width: '75%',
    height: 50,
    borderRadius: 10,
    overflow: 'hidden',
    marginTop: 10,
    backgroundColor: '#08d4c4',
  },
  gradient: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  buttonText: {
    color: '#fff',
    fontSize: 16,
  },
  imageUpload: {
    width: 100,
    height: 100,
    borderRadius: 50,
    backgroundColor: '#ddd',
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: 20,
  },
  uploadIcon: {
    width: 100,
    height: 100,
  },
  smallUploadIcon: {
    width: 30,
    height: 30,
    position: 'absolute',
    right: 0,
    bottom: 0,
  },
  socialContainer: {
    alignItems: 'flex-start',
    marginLeft: width * 0.05,
  },
  socialLabel: {
    fontWeight: 'bold',
    fontSize: 14,
    justifyContent: 'flex-end',
  },
  descriptionContainer: {
    alignItems: 'center',
  },
  buttonContainer: {
    alignItems: 'center',
    justifyContent: 'center',
    marginTop: 20,
  },
});
