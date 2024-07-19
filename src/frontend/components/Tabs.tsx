// // components/Tabs.js

// <Tabs step={step} setStep={setStep} />

// {step === 0 && (
//   <View style={styles.dateContainer}>
//     <Text style={{textAlign: 'left', fontWeight: 'bold', color: '#000', fontSize: 14}}>Selecione a data da atividade:</Text>
//     <DateInput 
//       placeholder="DD/MM/YYYY"
//       value={date}
//       onChangeText={setDate}
//     />
//   </View>
// )}

// {step === 1 && (
//   <View style={styles.titleContainer}>
//     <InputComponent
//       type="input"
//       label="Dê um título para o registro"
//       placeholder={"Digite o título..."}
//       maxLength={40}
//     />
//   </View>
// )}

// {step === 2 && (
//   <View style={styles.descriptionContainer}>
//     <InputComponent
//       type="textarea"
//       label="Adicione uma descrição"
//       placeholder={"Digite a descrição..."}
//       maxLength={2000}
//     />
//   </View>
// )}


import React from 'react';
import { View, Text, TouchableOpacity, StyleSheet, Dimensions } from 'react-native';

const window = Dimensions.get('window');

const Tabs = ({ step, setStep, page }) => {
  
  return (
    <View style={styles.tabsContainer}>
      <TouchableOpacity onPress={() => setStep(0)} style={styles.tab}>
        <Text style={[styles.tabText, step === 0 ? styles.activeTabText : styles.inactiveTabText]}>{page == 'Activity' ? 'Data' : 'Seu projeto'}</Text>
        <View style={[styles.lineTab, step === 0 ? styles.activeLineTab : styles.inactiveLineTab]} />
      </TouchableOpacity>
      <TouchableOpacity onPress={() => setStep(1)} style={styles.tab}>
        <Text style={[styles.tabText, step === 1 ? styles.activeTabText : styles.inactiveTabText]}>{page == 'Activity' ? 'Título' : 'Interesse'}</Text>
        <View style={[styles.lineTab, step === 1 ? styles.activeLineTab : styles.inactiveLineTab]} />
      </TouchableOpacity>
      <TouchableOpacity onPress={() => setStep(2)} style={styles.tab}>
        <Text style={[styles.tabText, step === 2 ? styles.activeTabText : styles.inactiveTabText]}>{page == 'Activity' ? 'Conteúdo' : 'Comentário'}</Text>
        <View style={[styles.lineTab, step === 2 ? styles.activeLineTab : styles.inactiveLineTab]} />
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  tabsContainer: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    marginVertical: 20,
    width: '100%',
    position: 'relative',
  },
  tab: {
    flex: 1,
    alignItems: 'center',
    marginBottom: 30,
  },
  tabText: {
    fontSize: 14,
  },
  activeTabText: {
    color: '#3A8A88',
    fontWeight: 'bold',
  },
  inactiveTabText: {
    color: '#999',
  },
  lineTab: {
    width: '90%',
    height: 6,
    borderRadius: 8,
    marginTop: 4,
  },
  activeLineTab: {
    backgroundColor: '#3A8A88',
  },
  inactiveLineTab: {
    backgroundColor: '#999',
  },
});

export default Tabs;
