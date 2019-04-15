/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "objects.hpp" file. It contains all the main objects used by
 * the Scoops Virtual Machine (SVM) and probably various other modules.
 *
 * You must remember that the only data type that exists in the Scoops 
 * Programming Language is the "ScpObj" class and all other data types are
 * merely its subclasses that inherit its rather humble functionality and
 * expand on it. 
 *
 * Such architecture allows SVM to have a single data stack
 * that deals with all the values it wishes to deal with. If there were
 * multiple classes, it would be impossible to do that as C++ std::stack
 * can only have values of the same type pushed to it..
 *
 */



#pragma once

#include <string>
#include <vector>



class ScpObj 
{ 
    std::string type;

    public:
    void set_type(std::string type) { this->type = type; }
    std::string get_type() { return this->type; }
};


class ScpBln : public ScpObj
{
    bool value;

    public:
    ScpBln (bool value)
    {
        this->set_type("bln");
        this->value = value;
    }
    
    bool get_value() { return this->value; }
};


class ScpChr : public ScpObj
{
    char value;

    public:
    ScpChr (char value)
    {
        this->set_type("chr");
        this->value = value;
    }

    char get_value() { return this->value; }
};


class ScpNum : public ScpObj
{
    double value;

    public:
    ScpNum (double value)
    {
        this->set_type("num");
        this->value = value;
    }

    double get_value() { return this->value; }
};


class ScpStr : public ScpObj
{
    std::vector<ScpChr> value;

    public:
    ScpStr () { this->set_type("str"); }

    unsigned int get_length() { return this->value.size(); }

    std::string get_value()
    {
        std::string value;
        for (ScpChr chr : this->value)
            value += chr.get_value();
        return value;
    }
    
    void add_chr(ScpChr chr) { this->value.push_back( chr.get_value() ); }

    void concat(ScpStr appendix) 
    {
        for ( ScpChr chr : appendix.get_value() )
            this->add_chr(chr);
    }
};

