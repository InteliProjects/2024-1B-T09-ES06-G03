import React from 'react';
import { View, Text, Dimensions, StyleSheet } from 'react-native';
import CeoAvatar from '../components/CeoAvatar'; // Certifique-se de importar corretamente
import Category from '../components/Category'; // Certifique-se de importar corretamente

const window = Dimensions.get('window');

const CardSinergyProject = ({ project, name, theme, category }) => {
  name = 'CEO ' + name;
  
  return (
    <View style={styles.container}>
      <View style={styles.column}>
        <View style={styles.row}>
          <View style={styles.rowContent}>
            <CeoAvatar size='w-12 h-12' link=''/>
            <View style={styles.textContainer}>
              <Text style={styles.projectText}>{project}</Text>
              <Text style={styles.nameText}>{name}</Text>
              <Text style={styles.themeText}>{theme}</Text>
            </View>
          </View>
          <View style={styles.categoryContainer}>
            <Category category={category} circleSize={'w-[32px] h-[32px]'} iconSize={20} />
          </View>
        </View>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingHorizontal: 20,
    paddingVertical: 10,
    backgroundColor: '#F6F6F6',
    borderRadius: 20,
    width: window.width * 0.9,
    borderTopRightRadius: 10,
    borderBottomRightRadius: 10,
    marginTop: window.height * 0.01,
    marginBottom: window.height * 0.01,
    shadowColor: 'rgba(0, 0, 0, 0.60)',
    shadowOpacity: 0.25,
    shadowOffset: { width: 0, height: 4 },
    elevation: 4,
    shadowRadius: 8
  },
  column: {
    flexDirection: 'column',
    flex: 1,
    justifyContent: 'center',
  },
  row: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  rowContent: {
    flexDirection: 'row',
    alignItems: 'center',
    
  },
  textContainer: {
    marginLeft: 10,
  },
  projectText: {
    fontSize: 18,
    fontWeight: '500',
    width: window.width * 0.6,
  },
  nameText: {
    fontSize: 14,
    fontWeight: '600',
    color: '#787878',
  },
  themeText: {
    fontSize: 14,
  },
  categoryContainer: {
    justifyContent: 'center',
  },
});

export default CardSinergyProject;
